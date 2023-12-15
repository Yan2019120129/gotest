package main

import (
	"fmt"
	"gotest/my_frame/config/gorm/database"
	"gotest/my_frame/models"
	"gotest/test"
)

func main() {
	var userInfo []*models.AdminUser
	database.DB.Find(userInfo)
	fmt.Println("data:", userInfo)
}

func init() {
	fmt.Println(test.Success)
}
