package main

import (
	"github.com/gin-gonic/gin"
	"my-frame/app/admin/router"
	"my-frame/config"
	"my-frame/module/logs"
)

func main() {

	// 配置gin
	cfg := config.GetGin()

	engin := gin.Default()

	// 初始化后台路由
	router.InitRouter(engin)

	// 初始化websocket
	//_ = okx.OkxInstance.ConnectWS()

	if err := engin.Run(cfg.Port); err != nil {
		logs.Logger.Error(err.Error())
	}
}