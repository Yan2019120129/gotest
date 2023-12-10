package database

import (
	"gotest/my_frame/config/database/mysql"
	"gotest/my_frame/config/database/postgresql"
	"gotest/my_frame/models"
)

func Init(cfg *models.Database) {
	switch cfg.UseDatabase {
	case models.DatabaseTypePostgresql:
		// 初始化Postgresql数据库
		postgresql.Init(&cfg.Postgresql)

	case models.DatabaseTypeMysql:
		// 初始化mysql数据库
		mysql.Init(&cfg.Mysql)
	}
}
