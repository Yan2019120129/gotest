package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gotest/my_frame/config"
)

type Mysql struct{}

func (m *Mysql) Connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(m.GetDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func (m *Mysql) GetDsn() string {
	cfg := &config.Cfg.Database.Mysql
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName)
}
