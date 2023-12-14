package test

import (
	"context"
	"fmt"
	esearch "gotest/my_frame/config/elasticsearch"
	"testing"
)

// CreateIndex 创建缩影
func TestCreateIndex(t *testing.T) {
	do, err := esearch.ES.Clint.Indices.Create("shopping").Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result", do.Index)
	fmt.Println("result：", do.Acknowledged)
	fmt.Println("result：", do.ShardsAcknowledged)
}

// DeleteIndex 删除索引
func TestDeleteIndex(t *testing.T) {
	do, err := esearch.ES.Clint.Indices.Delete("shopping").Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", do.Shards_)
	fmt.Println("result:", do.Acknowledged)
}

// FindIndexAll 获取节点信息
func TestFindIndexAll(t *testing.T) {
	do, err := esearch.ES.Clint.Indices.GetTemplate().Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", do)
}

// InsertDocument 插入文档
func TestInsertDocument(t *testing.T) {

}
