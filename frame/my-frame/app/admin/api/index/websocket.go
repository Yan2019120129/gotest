package index

import (
	"github.com/gin-gonic/gin"
	"my-frame/app/admin/service/dto"
)

// WebsocketServer  websocket服务
func WebsocketServer(c *gin.Context) {

	c.JSON(dto.Success(""))
}
