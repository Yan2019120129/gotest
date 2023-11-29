package init_config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Postgresql struct{}

func (m *Postgresql) Connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(m.GetDsn()), &gorm.Config{})
	if err != nil {
		log.Println("连接Mysql出错！！！")
		panic(err)
	}
	return db
}

func (m *Postgresql) GetDsn() string {
	cfg := &Cfg.Database.Postgresql
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Pass, cfg.DbName, cfg.Port)
}
