package main

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/api/admin"
	"gotest/my_frame/config"
	"net/http"
	"time"
)

func main() {
	// 初始化配置
	config.InitDatabase()

	// 配置gin
	router := gin.Default()
	s := &http.Server{
		Addr:           config.Cfg.Gin.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(config.Cfg.Gin.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Cfg.Gin.WriteTimeout) * time.Second,
		MaxHeaderBytes: config.Cfg.Gin.MaxHeaderBytes,
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
