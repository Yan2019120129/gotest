package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/my_frame/config/mysql"
	"gotest/my_frame/config/redis"
	"gotest/my_frame/models"
)

func Redis(c *gin.Context) {
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

func Mysql(c *gin.Context) {
	userInfo := new(models.User)
	rep := mysql.Db.Model(userInfo).Where("id=?", 1).Find(userInfo)
	if rep.Error != nil {
		panic(rep.Error)
	}
	c.JSON(200, userInfo)
}
