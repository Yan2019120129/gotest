package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"my-frame/app/admin/router"
	"my-frame/configs"
	"my-frame/module/logs"
)

func main() {
	InitApp()
}

// InitApp 初始化项目
func InitApp() {
	// 配置gin
	cfg := configs.GetGin()

	engin := gin.Default()

	// 初始化后台路由
	router.InitRouter(engin)

	zap.ReplaceGlobals(logs.Logger)

	if err := engin.Run(cfg.Port); err != nil {
		zap.L().Error(err.Error())
	}
}
