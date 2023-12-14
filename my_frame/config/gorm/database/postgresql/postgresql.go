package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest/my_frame/config"
)

var open gorm.Dialector

// Init 初始化Postgresql
func Init() {
	cfg := config.GetPostgres()
	open = postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Pass, cfg.DbName, cfg.Port))
}

// GetOpen 获取链接信息
func GetOpen() gorm.Dialector {
	return open
}
