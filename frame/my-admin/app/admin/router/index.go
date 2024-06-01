package admin

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	r.Static("/website/public/uploads", "./website/public/uploads")
}
