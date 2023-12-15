package index

import (
	"github.com/gin-gonic/gin"
	"gotest/my_frame/app/admin/service/index"
)

func Index(c *gin.Context) {
	data, err := index.Index()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, data)
}
