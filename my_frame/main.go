package main

import (
	"gotest/my_frame/config/gin"
)

func main() {

	//// 初始化配置文件，全局依赖配置文件配置
	//config.Init()
	//
	//// 初始化配置
	//database.Init()

	// 初始化redis
	//redis.Init()

	// 初始化Elasticsearch
	//esearch.Init()

	//// 初始化gin
	gin.Init()

}
