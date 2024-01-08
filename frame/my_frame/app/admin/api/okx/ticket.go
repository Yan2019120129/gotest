package okx

import (
	"github.com/gin-gonic/gin"
	"gotest/frame/my_frame/app/admin/service/okx"
)

// TickerIndex 获取行情数据
func TickerIndex(c *gin.Context) {

	okxserver.TickerIndex(c.Writer, c.Request)
}
