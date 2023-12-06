package postgresql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest/my_frame/models"
	"log"
)

var Db *gorm.DB

// GetDsn 获取链接信息
func GetDsn(cfg *models.DatabaseConfig) string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Pass, cfg.DbName, cfg.Port)
}

// Init 初始化Postgresql
func Init(cfg *models.DatabaseConfig) {
	var err error
	Db, err = gorm.Open(mysql.Open(GetDsn(cfg)), &gorm.Config{})
	if err != nil {
		log.Println("连接Mysql出错！！！")
		panic(err)
	}

}
