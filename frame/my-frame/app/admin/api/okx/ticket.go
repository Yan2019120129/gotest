package okx

import (
	"github.com/gin-gonic/gin"
	okxserver "my-frame/app/admin/service/okx"
)

// TickerIndex 获取行情数据
func TickerIndex(c *gin.Context) {

	okxserver.TickerIndex(c.Writer, c.Request)
}
