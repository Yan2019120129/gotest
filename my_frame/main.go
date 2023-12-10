package main

import (
	"gotest/my_frame/config"
	esearch "gotest/my_frame/config/elasticsearch"
	"gotest/my_frame/test"
)

func main() {

	// 初始化配置文件，全局依赖配置文件配置
	cfg := config.GetConfig()

	// 初始化配置
	//database.Init(&cfg.Database)

	// 初始化redis
	//redis.Init(&cfg.Redis)

	// 初始化Elasticsearch
	esearch.Init(&cfg.Elasticsearch)

	//// 初始化gin
	//gin.Init(&cfg.Gin)

	test.DeleteIndex("customer")
}
