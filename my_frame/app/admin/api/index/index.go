package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/my_frame/config/redis"
)

func Index(c *gin.Context) {
	rds := redis.Rds.Get()

	defer rds.Close()
	_, err := rds.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := rds.Do("Get", "abc")
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)

	c.JSON(200, r)
}
