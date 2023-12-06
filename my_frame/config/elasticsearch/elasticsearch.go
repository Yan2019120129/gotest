package esearch

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gotest/my_frame/models"
	"log"
	"strings"
)

var Clinet *elasticsearch.Client

type ESearch struct {
}

// Init 初始化elasticsearch
func Init(cfg *models.Elasticsearch) {

	// 添加配置，初始化连接
	var err error
	Clinet, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.Url},
	})
	if err != nil {
		panic(err)
	}
}

// Insert 插入数据
func (es *ESearch) Insert(index, documentId string, data []byte) {
	req := esapi.IndexRequest{
		Index:      index,
		Body:       strings.NewReader(string(data)),
		DocumentID: documentId,
		Refresh:    "true",
	}
	es.IndexRequest(&req)
}

// SearchAll 查询全部文档
func (es *ESearch) SearchAll(index []string, searchParams interface{}) interface{} {
	req := esapi.SearchRequest{
		Index: index,
		Body:  es.ConversionReader(searchParams),
	}
	return es.SearchRequest(&req)
}

// ConversionReader 转换为reader格式数据
func (es *ESearch) ConversionReader(data interface{}) *strings.Reader {
	bt, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(bt))
}

// SearchRequest 发送查找请求
func (es *ESearch) SearchRequest(request *esapi.SearchRequest) interface{} {
	res, err := request.Do(context.Background(), Clinet)
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

// IndexRequest 发送查找请求
func (es *ESearch) IndexRequest(request *esapi.IndexRequest) {
	res, err := request.Do(context.Background(), Clinet)
	if err != nil {
		log.Fatalf("Error performing the search request: %s", err)
	}
	defer res.Body.Close()
}
