package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gotest/common/utils"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// ================== DB ==================

var db *sql.DB

func initDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/craw?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

// ================== 通用模型 ==================

type Product struct {
	SourceSite string
	SourceID   int64
	Name       string
	Price      float64
	Image      string
}

// ================== WTAPS API结构 ==================

type WTAPSResponse struct {
	Products []WTAPSProduct `json:"products"`
}

type WTAPSProduct struct {
	ID       int64          `json:"id"`
	Title    string         `json:"title"`
	Images   []WTAPSImage   `json:"images"`
	Variants []WTAPSVariant `json:"variants"`
}

type WTAPSImage struct {
	Src string `json:"src"`
}

type WTAPSVariant struct {
	Price string `json:"price"`
}

// ================== 站点定义 ==================

type Site struct {
	Name  string
	Fetch func() ([]Product, error)
}

type Result struct {
	Site  string
	Count int
	Err   error
}

// ================== WTAPS 爬取 ==================

const (
	WTAPSBaseURL = "https://wtaps.com/en/collections/all/products.json"
)

func fetchWTAPS() ([]Product, error) {
	var resultAll []Product

	h := utils.NewHttp()
	req, err := h.Get(WTAPSBaseURL)
	if err != nil {
		return nil, err
	}

	var res WTAPSResponse
	err = json.Unmarshal(req, &res)
	if err != nil {
		return nil, err
	}

	if len(res.Products) == 0 {
		return nil, errors.New("no products found")
	}

	for _, p := range res.Products {
		resultAll = append(resultAll, convertWTAPS(p))
	}

	return resultAll, nil
}

func convertWTAPS(p WTAPSProduct) Product {
	var img string
	if len(p.Images) > 0 {
		img = p.Images[0].Src
	}

	var minPrice float64
	first := true

	for _, v := range p.Variants {
		var price float64
		fmt.Sscanf(v.Price, "%f", &price)

		if first || price < minPrice {
			minPrice = price
			first = false
		}
	}

	return Product{
		SourceSite: "wtaps.com",
		SourceID:   p.ID,
		Name:       p.Title,
		Price:      minPrice,
		Image:      img,
	}
}

// ================== 其他站点（占位） ==================

func fetchStoneIsland() ([]Product, error) {
	return nil, fmt.Errorf("stoneisland not implemented")
}

func fetchBAPE() ([]Product, error) {
	return nil, fmt.Errorf("bape not implemented")
}

func fetchSupreme() ([]Product, error) {
	return nil, fmt.Errorf("supreme not implemented")
}

// ================== 入库 ==================

func saveProducts(products []Product) error {
	sqlStr := `
	INSERT INTO craw_products (source_id, source_site, name, price, main_image_url)
	VALUES (?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
	name = VALUES(name),
	price = VALUES(price),
	main_image_url = VALUES(main_image_url)
	`

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range products {
		if p.Name == "" || p.SourceID == 0 {
			continue
		}

		_, err := stmt.Exec(p.SourceID, p.SourceSite, p.Name, p.Price, p.Image)
		if err != nil {
			log.Println("DB error:", err)
		}
	}

	return nil
}

// ================== 站点注册 ==================

func getSites() []Site {
	return []Site{
		{
			Name:  "wtaps.com",
			Fetch: fetchWTAPS,
		},
		{
			Name:  "stoneisland.com",
			Fetch: fetchStoneIsland,
		},
		{
			Name:  "bape.com",
			Fetch: fetchBAPE,
		},
		{
			Name:  "supreme",
			Fetch: fetchSupreme,
		},
	}
}

// ================== 并发执行 ==================

func runAllSites() []Result {
	sites := getSites()

	var wg sync.WaitGroup
	resultCh := make(chan Result, len(sites))

	for _, s := range sites {
		wg.Add(1)

		go func(site Site) {
			defer wg.Done()

			products, err := site.Fetch()

			if err != nil {
				resultCh <- Result{
					Site:  site.Name,
					Err:   err,
					Count: 0,
				}
				return
			}

			// 入库
			if len(products) > 0 {
				for _, product := range products {
					fmt.Println(product)
				}
				//_ = saveProducts(products)
				fmt.Println("len(products):", len(products))
			}

			resultCh <- Result{
				Site:  site.Name,
				Count: len(products),
				Err:   nil,
			}
		}(s)
	}

	wg.Wait()
	close(resultCh)

	var results []Result
	for r := range resultCh {
		results = append(results, r)
	}

	return results
}

// ================== main ==================

func main() {
	log.Println("🚀 Multi-site crawler started")

	initDB()

	results := runAllSites()

	success := 0
	fail := 0

	for _, r := range results {
		if r.Err != nil {
			log.Printf("❌ [%s] failed: %v", r.Site, r.Err)
			fail++
		} else {
			log.Printf("✅ [%s] success, %d items", r.Site, r.Count)
			success++
		}
	}

	log.Printf("🎯 DONE: success=%d fail=%d", success, fail)
}
