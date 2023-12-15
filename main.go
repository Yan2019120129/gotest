package main

import (
	"fmt"
	"gotest/my_frame/config/gorm/database"
	"gotest/my_frame/models"
)

func main() {
	userInfo := &models.User{}
	result := database.DB.First(userInfo)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println("data:", userInfo)
}
