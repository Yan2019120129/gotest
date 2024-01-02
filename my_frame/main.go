package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/router"
	"gotest/my_frame/config"
	"gotest/my_frame/module/okx"
)

func main() {

	// 配置gin
	cfg := config.GetGin()
	engin := gin.Default()

	// 初始化后台路由
	router.InitRouter(engin)

	// 初始化websocket
	_ = okx.OkxInstance.ConnectWS()

	if err := endless.ListenAndServe(cfg.Port, engin); err != nil {
		panic(err)
	}
}
