package http_t

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/rand"
	"net/http"
	"time"
)

// Product 数据库模型属性
type Product struct {
	Id            int64   //主键
	ParentId      int64   //上级ID
	AdminId       int64   //管理员ID
	CategoryId    int64   //类目ID
	ShopId        int64   //店铺ID 0批发中心
	AssetsId      int64   //资产ID
	Name          string  //标题
	Images        string  //图片列表
	OriginalMoney float64 //原价
	Money         float64 //现价
	Increase      float64 //涨幅
	Type          int64   //类型 1自营商品 2批发商品
	Sort          int64   //排序
	Status        int64   //状态 -2删除 -1禁用 10启用
	Sales         int64   //销售量
	Rating        float64 //评分
	Nums          int64   //限购 -1 无限制
	Used          int64   //已用
	Total         int64   //总数
	Data          string  //数据
	Describes     string  //描述
	UpdatedAt     int64   //更新时间
	CreatedAt     int64   //创建时间
}

type copyParams struct {
	ParentId int64   `json:"parent_id" validate:"required"`
	ShopId   int64   `json:"shop_id" validate:"required"`
	Increase float64 `json:"increase" validate:"required,gt=0"`
}

func readData() ([]*copyParams, error) {
	dsn := "root:Aa123098..@tcp(192.168.5.15:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	// Set the table name (optional, GORM will pluralize the struct name by default)
	db.Table("product")

	// Retrieve all records
	products := make([]*Product, 0)
	result := db.Find(&products)

	params := make([]*copyParams, 0)
	for _, v := range products {
		// 设置随机数种子，以确保每次运行都能得到不同的结果
		rand.Seed(time.Now().UnixNano())

		// 生成1到5之间的随机整数
		randomNumber := rand.Intn(5) + 1
		param := new(copyParams)
		// 生成20到50之间的随机浮点数
		randomFloat := rand.Float64()*(50-20) + 20

		param.ParentId = v.Id
		param.ShopId = int64(randomNumber)
		param.Increase = randomFloat
		params = append(params, param)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("Error retrieving data: %w", result.Error)
	}
	return params, nil
}

func NewHttp() {
	products, err := readData()
	//fmt.Println("result:", products)
	// 准备 POST 请求的数据
	url := "http://192.168.5.15:8009/product/index/copy"
	//params := new(copyParams)
	//params.ParentId = 800
	//params.ShopId = 10
	//params.Increase = 10.50
	payload, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}

	// 创建一个字节缓冲区，并将数据写入其中
	buffer := bytes.NewBuffer(payload)

	// 发送 POST 请求
	response, err := http.Post(url, "application/json", buffer)
	if err != nil {
		fmt.Println("HTTP POST request failed:", err)
		return
	}
	defer response.Body.Close()

	// 处理响应
	fmt.Println("HTTP Status Code:", response.Status)

	// 读取响应体
	body := new(bytes.Buffer)
	_, err = body.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Response Body:", body.String())
}
