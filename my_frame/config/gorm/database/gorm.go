package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
)

import (
	"gotest/my_frame/config"
	"gotest/my_frame/config/gorm/database/mysql"
	"gotest/my_frame/config/gorm/database/postgresql"
	"gotest/my_frame/models"
)

// 定义once 保证初始化只执行一次
var _once sync.Once

// DB 定义全局数据库对象
var DB *gorm.DB

func init() {
	if DB == nil {
		_once.Do(func() {
			var err error
			cfg := config.GetGorm()
			if DB, err = gorm.Open(getDatabaseOpen(cfg.UseDatabase), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{ // 命名策略
					SingularTable: cfg.SingularTable, // 单表去复数s
				},
				QueryFields: cfg.QueryFields,                     // 是否全字段映射
				Logger:      logger.Default.LogMode(logger.Info), // 日志级别
			}); err != nil {
				panic(err)
			}
			fmt.Printf("内存地址：%p----->Gorm.DB实例创建成功！！！\n", DB)
		})
	} else {
		fmt.Println("已经存在DB实例")
	}
}

// DatabaseConfig 获取数据库配置
func getDatabaseOpen(useDatabase string) (databaseOpen gorm.Dialector) {
	// 选用数据库
	switch useDatabase {
	case models.DatabaseTypePostgresql:
		// 初始化Postgresql数据库
		databaseOpen = postgresql.GetOpen()

	case models.DatabaseTypeMysql:
		// 初始化mysql数据库
		databaseOpen = mysql.GetOpen()
	}
	return
}
