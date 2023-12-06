package config

import (
	"gorm.io/gorm"
	esearch "gotest/my_frame/config/elasticsearch"
	"gotest/my_frame/config/mysql"
	"gotest/my_frame/config/postgresql"
	"gotest/my_frame/config/redis"
)

var Db *gorm.DB

func InitDatabase() {
	// 初始化配置文件，全局依赖配置文件配置
	InitConfig()

	switch Cfg.Database.UseDatabase {
	// 初始化Postgresql数据库
	case DatabaseTypePostgresql:
		postgresql := new(postgresql.Postgresql)
		Db = postgresql.Connect()
	// 初始化mysql数据库
	case DatabaseTypeMysql:
		Db = new(mysql.Mysql).Connect()
	}

	// 初始化redis
	redis.InitRedis()

	// 初始化Elasticsearch
	esearch.InitElasticsearch()
}
