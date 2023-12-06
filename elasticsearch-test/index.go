package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

// User 用户表
type User struct {
	Id          int     `fake:"-"`
	AdminId     int     `fake:"{randomint:[1]}"`
	ParentId    int     `fake:"{randomint:[1]}"`
	UserName    string  `fake:"{name}"`
	NickName    string  `fake:"{name}"`
	Email       string  `fake:"{email}"`
	Telephone   string  `fake:"{phone}"`
	Avatar      string  `fake:"{imageurl}"`
	Sex         int     `fake:"{randomint:[1,2]}"`
	Birthday    int     `fake:"{number:946656000,1104508800}"`
	Password    string  `fake:"{randomstring:['1234567890','qwertyuiopasdfghjklzxcvbnm','QWERTYUIOPASDFGHJKLZXCVBNM']}"`
	SecurityKey string  `fake:"{randomstring:['1234567890','qwertyuiopasdfghjklzxcvbnm','QWERTYUIOPASDFGHJKLZXCVBNM']}"`
	Money       float64 `fake:"{float64range:100,1000}"`
	Type        int     `fake:"{randomint:[-2,-1,10]}"`
	Status      int     `fake:"{randomint:[-2,-1,10]}"`
	Data        string  `fake:"{loremipsumsentence}"`
	UpdatedAt   int     `fake:"{day}"`
	CreatedAt   int     `fake:"{day}"`
}

const url = "http://47.101.70.217:1014"

var es *elasticsearch.Client

var config = elasticsearch.Config{
	Addresses: []string{url},
}

func main() {
	//account := models.Account{
	//	AccountNumber: gofakeit.Number(1, 100000),
	//	Address:       gofakeit.Address(),
	//	Age:           gofakeit.Number(18, 25),
	//	Balance:       gofakeit.Number(500, 2000),
	//	City:          gofakeit.City(),
	//	Email:         gofakeit.Email(),
	//	Employer:      gofakeit.Name(),
	//	Firstname:     gofakeit.FirstName(),
	//	Gender:        gofakeit.Gender(),
	//	Lastname:      gofakeit.LastName(),
	//	State:         gofakeit.State(),
	//}
	//data, err := json.Marshal(account)
	//if err != nil {
	//	return
	//}

	//insert("test", "", data)
	searchBody := map[string]interface{}{"query": map[string]interface{}{"match_all": map[string]interface{}{}}}
	data := searchAll([]string{"test"}, searchBody)
	fmt.Println("data:", data)
}

// 插入数据
func insert(index, documentId string, data []byte) {
	req := esapi.IndexRequest{
		Index:      index,                           // Index name
		Body:       strings.NewReader(string(data)), // Document body
		DocumentID: documentId,                      // Document ID
		Refresh:    "true",                          // Refresh
	}
	indexRequest(&req)
}

func searchAll(index []string, searchParams interface{}) interface{} {
	req := esapi.SearchRequest{
		Index: index,
		Body:  conversionReader(searchParams),
	}
	return searchRequest(&req)
}
func conversionReader(data interface{}) *strings.Reader {
	bt, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(bt))
}

func searchRequest(request *esapi.SearchRequest) interface{} {
	// 执行搜索请求
	res, err := request.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error performing the search request: %s", err)
	}

	defer res.Body.Close()

	// 解析响应
	var response map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response: %s", err)
	}

	return response
}

func indexRequest(request *esapi.IndexRequest) {
	// 执行搜索请求
	res, err := request.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error performing the search request: %s", err)
	}
	defer res.Body.Close()
}

func init() {
	var err error
	es, err = elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
}
