package index

import (
	"github.com/gin-gonic/gin"
	"gotest/frame/my_frame/app/admin/service/index"
)

// Login 登录接口
func Login(c *gin.Context) {
	data, err := indexserver.Login()
	if err != nil {
		return
	}
	c.JSON(200, data)
	return
}
