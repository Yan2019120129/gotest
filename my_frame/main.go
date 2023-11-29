package main

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/api/admin"
	"gotest/my_frame/database/redis"
)

func main() {
	// 配置gin
	ctx := gin.Default()

	// 初始化后台路由
	admin.InitRouter(ctx)

	// 初始化Redis
	redis.InitRedis()

	// 运行服务
	err := ctx.Run()
	if err != nil {
		return
	}
}
