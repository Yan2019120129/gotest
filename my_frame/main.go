package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/router"
	"gotest/my_frame/module"
)

func main() {
	// 配置gin
	cfg := module.GetGin()
	engin := gin.Default()

	// 初始化后台路由
	router.InitRouter(engin)

	if err := endless.ListenAndServe(cfg.Port, engin); err != nil {
		panic(err)
	}
}
