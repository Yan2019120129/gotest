package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gotest/common/config"
	"gotest/common/module/gorm/database/mysql"
	"gotest/common/module/gorm/database/postgresql"
	"gotest/common/module/logs"
	"sync"
)

// 定义公共数据库配置
var databaseOpen = map[string]gorm.Dialector{
	config.DatabaseTypeMysql:      mysql.GetOpen(),
	config.DatabaseTypePostgresql: postgresql.GetOpen(),
}

// 定义once 保证初始化只执行一次
var _once sync.Once

// DB 定义全局数据库对象
var DB *gorm.DB

// 初始化数据库连接，保证只执行一次
func init() {
	if DB == nil {
		_once.Do(func() {
			// 获取配置文件数据
			cfg := config.GetGorm()
			var err error
			if DB, err = gorm.Open(databaseOpen[cfg.UseDatabase], &gorm.Config{
				NamingStrategy: schema.NamingStrategy{ // 命名策略
					SingularTable: cfg.SingularTable, // 单表去复数s
				},
				QueryFields: cfg.QueryFields, // 是否全字段映射
				//Logger: logger.Default.LogMode(logger.Info), // 日志级别
				Logger: logs.Instance.LogMode(logs.LeverGorm[config.GetZap().Level]), // 日志级别
			}); err != nil {
				logs.Logger.Error("gorm", zap.String("method", "init"), zap.Error(err))
			}
			logs.Logger.Info("gorm", zap.String("method", "init"), zap.String("instance", fmt.Sprintf("%p", DB)))
		})
	} else {
		logs.Logger.Info("gorm", zap.String("method", "init"), zap.String("instance existed", fmt.Sprintf("%p", DB)))
	}
}
