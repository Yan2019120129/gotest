package main

import (
	"gotest/my_frame/config"
	"gotest/my_frame/config/database"
	"gotest/my_frame/test"
)

func main() {

	// 初始化配置文件，全局依赖配置文件配置
	cfg := config.GetConfig()

	// 初始化配置
	database.Init(&cfg.Database)

	// 初始化redis
	//redis.Init(&cfg.Redis)

	// 初始化Elasticsearch
	//esearch.Init(&cfg.Elasticsearch)

	//// 初始化gin
	//gin.Init(&cfg.Gin)

	//test.TestTime()
	//test.TestMap()

	//err := test.Synchronous([]*test.SynchronousParams{
	//	{Table: "wallet_assets", AdminId: 2, WhereField: []string{"name"}},
	//	{Table: "lang", AdminId: 2, WhereField: []string{"name", "alias"}},
	//	{Table: "article", AdminId: 2, WhereField: []string{"name", "type"}},
	//	{Table: "level", AdminId: 2, WhereField: []string{"name", "money"}},
	//	{Table: "real_name_auth", AdminId: 2, WhereField: []string{"real_name", "number"}},
	//	{Table: "translate", AdminId: 2, WhereField: []string{"name", "field", "value"}},
	//	{Table: "admin_setting", AdminId: 2, WhereField: []string{"name", "type", "field"}},
	//})
	//if err != nil {
	//	panic(err)
	//}

	err := test.Copy([]*test.CopyParams{
		{Table: "wallet_assets", AdminId: 2},
		{Table: "lang", AdminId: 2},
		{Table: "article", AdminId: 2},
		{Table: "level", AdminId: 2},
		{Table: "real_name_auth", AdminId: 2},
		{Table: "translate", AdminId: 2},
		{Table: "admin_setting", AdminId: 2},
	})
	if err != nil {
		panic(err)
	}

}
