package json_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/goccy/go-json"
	"gotest/middleware/mysql_test/dto"
	"time"
)

// TestJson 测试json解析
func TestJson() {
	data := &dto.ProductData{
		InstId:    gofakeit.MonthString(),
		Last:      gofakeit.Float64Range(1000, 10000),
		LastSz:    gofakeit.Float64Range(10, 100),
		Open24h:   gofakeit.Float64Range(500, 1000),
		High24h:   gofakeit.Float64Range(500, 1000),
		Low24h:    gofakeit.Float64Range(500, 1000),
		Vol24h:    gofakeit.Float64Range(500, 1000),
		Amount24h: gofakeit.Float64Range(500, 1000),
		Ts:        time.Now().Unix(),
	}
	//productDataBytes, err := json.Marshal(data)
	//if err != nil {
	//	panic(err)
	//}
	test := map[string]interface{}{
		"id":          1,
		"category_id": 2,
		"images":      []string{"/assets/currency/doge.png"},
		"name":        gofakeit.Name(),
		"money":       gofakeit.Float64Range(10, 100),
		"type":        1,
		"sales":       1,
		"nums":        1,
		"used":        1,
		"total":       1,
		"isCollect":   true,
		"data":        data,
	}
	dataBytes, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	product := dto.IndexData{}
	if err = json.Unmarshal(dataBytes, &product); err != nil {
		panic(err)
	}
	fmt.Println("data:", product)
}
