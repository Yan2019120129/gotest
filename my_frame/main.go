package main

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/api/admin"
	"gotest/my_frame/init_config"
)

func main() {
	// 初始化配置
	init_config.InitDatabase()

	// 配置gin
	ctx := gin.Default()

	// 初始化后台路由
	admin.InitRouter(ctx)

	// 运行服务
	err := ctx.Run()
	if err != nil {
		return
	}
}
