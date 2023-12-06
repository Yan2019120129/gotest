package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

func main() {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://47.101.70.217:1014"},
	})
	if err != nil {
		panic(err)
	}
	req := esapi.IndexRequest{
		Index:      "shopping",                              // Index name
		Body:       strings.NewReader(`{"title" : "Test"}`), // Document body
		DocumentID: "1",                                     // Document ID
		Refresh:    "true",                                  // Refresh
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}

//// 准备一个简单的搜索请求，匹配所有文档
//	var buf string
//	query := map[string]interface{}{
//		"query": map[string]interface{}{
//			"match_all": map[string]interface{}{},
//		},
//	}
//	if err := es.EncodeJSON(&buf, query); err != nil {
//		log.Fatalf("Error encoding query: %s", err)
//	}
//
//	// 执行搜索请求
//	res, err := es.Search(
//		es.Search.WithContext(context.Background()),
//		es.Search.WithIndex("shopping"), // 替换为你的索引
//		es.Search.WithBodyString(buf),
//		es.Search.WithTrackTotalHits(true),
//		es.Search.WithPretty(),
//	)
//	if err != nil {
//		log.Fatalf("Error performing the search request: %s", err)
//	}
//	defer res.Body.Close()
//
//	// 解析响应
//	var response map[string]interface{}
//	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
//		log.Fatalf("Error parsing the response: %s", err)
//	}
//
//	// 输出搜索结果
//	fmt.Printf("Search Results:\n%s\n", response)
