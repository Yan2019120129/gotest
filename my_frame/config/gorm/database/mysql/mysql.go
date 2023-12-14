package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest/my_frame/config"
)

var open gorm.Dialector

// Init 初始化mysql
func Init() {
	cfg := config.GetMysql()
	mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName))
}

func GetOpen() gorm.Dialector {
	return open
}
