package okx

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/service/dto"
	"gotest/my_frame/app/admin/service/okx"
)

// TickerIndex 获取行情数据
func TickerIndex(c *gin.Context) {
	params := &dto.TickerParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(dto.Error(err))
		return
	}

	data, err := okxserver.TickerIndex(c.Writer, c.Request, params)
	if err != nil {
		c.JSON(dto.Error(err))
		return
	}

	c.JSON(dto.Success(data))
}
