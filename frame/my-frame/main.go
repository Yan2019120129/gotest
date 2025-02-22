package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"my-frame/app/admin/router"
	"my-frame/config"
	"my-frame/module/logs"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	InitApp()
}

// InitApp 初始化项目
func InitApp() {
	// 配置gin
	cfg := config.GetGin()

	engin := gin.Default()

	// 初始化后台路由
	router.InitRouter(engin)

	zap.ReplaceGlobals(logs.Logger)

	if err := engin.Run(cfg.Port); err != nil {
		zap.L().Error(err.Error())
	}

	// 检测系统信号关闭
	// syscall.SIGINT 中断信号，通常在用户按下 Ctrl+C 时发送。
	// syscall.SIGTERM 终止信号，通常用于请求程序终止。
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
