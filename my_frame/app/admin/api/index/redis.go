package index

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/service/dto"
	"gotest/my_frame/app/admin/service/index"
)

// Rds 测试redis
func Rds(c *gin.Context) {
	params := &dto.RdsParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(dto.Error(err))
		return
	}
	if data, err := index.Rds(params); err != nil {
		c.JSON(dto.Error(err))
	} else {
		c.JSON(dto.Success(data))
	}

}

// SubRds 订阅信息
func SubRds(c *gin.Context) {
	params := &dto.RdsPublishAndParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(dto.Error(err))
		return
	}

	if data, err := index.SubRds(params); err != nil {
		c.JSON(dto.Error(err))
	} else {
		c.JSON(dto.Success(data))
	}
}

// Publish 发布消息
func Publish(c *gin.Context) {
	params := &dto.RdsPublishAndParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(dto.Error(err))
		return
	}

	if data, err := index.Publish(params); err != nil {
		c.JSON(dto.Error(err))
	} else {
		c.JSON(dto.Success(data))
	}
}
