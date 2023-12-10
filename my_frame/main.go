package main

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/api/admin"
	"gotest/my_frame/config"
	esearch "gotest/my_frame/config/elasticsearch"
	"gotest/my_frame/config/redis"
	"net/http"
	"time"
)

func main() {

	// 初始化配置文件，全局依赖配置文件配置
	cfg := config.GetConfig()

	// 初始化配置
	config.InitDatabase(&cfg.Database)

	// 初始化redis
	redis.Init(&cfg.Redis)

	// 初始化Elasticsearch
	esearch.Init(&cfg.Elasticsearch)

	// 配置gin
	router := gin.Default()
	s := &http.Server{
		Addr:           cfg.Gin.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.Gin.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.Gin.WriteTimeout) * time.Second,
		MaxHeaderBytes: cfg.Gin.MaxHeaderBytes,
	}

	// 初始化后台路由
	admin.InitRouter(router)

	// 运行服务
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
		return
	}
}
