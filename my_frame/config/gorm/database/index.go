package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gotest/my_frame/config"
	"gotest/my_frame/config/gorm/database/mysql"
	"gotest/my_frame/config/gorm/database/postgresql"
	"gotest/my_frame/models"
)

var Db *gorm.DB

func Init() {
	var err error
	cfg := config.GetGorm()
	Db, err = gorm.Open(getDatabaseOpen(cfg.UseDatabase), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ // 命名策略
			SingularTable: cfg.SingularTable, // 单表去复数s
		},
		QueryFields: cfg.QueryFields,                     // 是否全字段映射
		Logger:      logger.Default.LogMode(logger.Info), // 日志级别
	})
	if err != nil {
		panic(err)
	}
}

// DatabaseConfig 获取数据库配置
func getDatabaseOpen(useDatabase string) (databaseOpen gorm.Dialector) {
	// 判断是有的是什么数据库
	switch useDatabase {
	case models.DatabaseTypePostgresql:

		// 初始化Postgresql数据库
		postgresql.Init()
		databaseOpen = postgresql.GetOpen()

	case models.DatabaseTypeMysql:

		// 初始化mysql数据库
		mysql.Init()
		databaseOpen = mysql.GetOpen()
	}
	return
}
