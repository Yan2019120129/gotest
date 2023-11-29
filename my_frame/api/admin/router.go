package admin

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/service/admin/index"
	"gotest/my_frame/service/admin/test"
)

func InitRouter(ctx *gin.Engine) {
	router := ctx.Group("/v1")
	{
		router.POST("/login")
		router.GET("/index", index.Index)
		router.GET("/test/redis", test.Redis)
		router.GET("/test/mysql", test.Mysql)
	}
}
