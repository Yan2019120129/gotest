package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _open gorm.Dialector

// init 初始化mysql
func init() {
	cfg := configs.GetMysql()
	_open = mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName))
}

func GetOpen() gorm.Dialector {
	return _open
}
