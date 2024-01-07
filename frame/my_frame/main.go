package main

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/router"
	"gotest/my_frame/config"
	"gotest/my_frame/module/logger"
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
		logger.Logger.Error(err.Error())
	}
}
