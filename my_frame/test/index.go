package test

import (
	"fmt"
	"golang.org/x/net/context"
	"gotest/my_frame/config/database/mysql"
	esearch "gotest/my_frame/config/elasticsearch"
	"gotest/my_frame/models"
	"log"
)

// CreateTable 创建表
func CreateTable() {
	result := mysql.Db.AutoMigrate(&models.Product{}, &models.ProductOrder{}, &models.ShopOrder{})
	if result.Error() != "" {
		log.Panicln(result.Error())
	}
}

// InsertMysql 向表添加数据
func InsertMysql() {

}

// CreateIndex 创建缩影
func CreateIndex(index string) {
	do, err := esearch.ES.Clint.Indices.Create(index).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result", do.Index)
	fmt.Println("result：", do.Acknowledged)
	fmt.Println("result：", do.ShardsAcknowledged)
}

// DeleteIndex 删除索引
func DeleteIndex(index string) {
	do, err := esearch.ES.Clint.Indices.Delete(index).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", do.Shards_)
	fmt.Println("result:", do.Acknowledged)
}

// FindIndexAll 获取节点信息
func FindIndexAll() {
	do, err := esearch.ES.Clint.Indices.GetTemplate().Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", do)
}

// InsertDocument 插入文档
func InsertDocument() {

}
