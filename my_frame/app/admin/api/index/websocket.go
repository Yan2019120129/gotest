package index

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/service/dto"
	"gotest/my_frame/app/admin/service/index"
)

// WebsocketServer  websocket服务
func WebsocketServer(c *gin.Context) {
	data, err := indexserver.WebsocketServer(c.Writer, c.Request)
	if err != nil {
		c.JSON(dto.Error(err))
		return
	}
	c.JSON(dto.Success(data))
}
