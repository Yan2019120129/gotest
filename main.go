package main

import (
	"fmt"
	"gotest/my_frame/config/gorm/database"
	"gotest/my_frame/models"
)

func main() {
	//DB, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", "root", "Aa123098..", "127.0.0.1", 3306, "basic")), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{ // 命名策略
	//		SingularTable: true, // 单表去复数s
	//	},
	//	QueryFields: false,                               // 是否全字段映射
	//	Logger:      logger.Default.LogMode(logger.Info), // 日志级别
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//db := database.DB
	//fmt.Println("db:", DB)
	//fmt.Println("db:", db)

	userInfo := &models.User{}
	result := database.DB.First(userInfo)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println("data:", userInfo)
}
