package router

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/api/index"
)

func InitRouter(ctx *gin.Engine) {
	router := ctx.Group("/v1")
	{
		router.GET("/login", index.Login)
		router.GET("/index", index.Index)
		router.GET("/init", index.Init)
		router.GET("/registration", index.Registration)
		router.GET("/websocket/server", index.WebsocketServer)
		router.POST("/redis/rds", index.Rds)
		router.POST("/redis/sub", index.SubRds)
		router.POST("/redis/publish", index.Publish)
	}
}
