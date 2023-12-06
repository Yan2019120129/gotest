package config

import (
	"gopkg.in/yaml.v3"
	esearch "gotest/my_frame/config/elasticsearch"
	"gotest/my_frame/config/mysql"
	"gotest/my_frame/config/postgresql"
	"gotest/my_frame/config/redis"
	"gotest/my_frame/models"
	"os"
)

func InitDatabase(cfg *models.Database) {

	switch cfg.UseDatabase {
	// 初始化Postgresql数据库
	case models.DatabaseTypePostgresql:
		postgresql.Init(&cfg.Postgresql)
	// 初始化mysql数据库
	case models.DatabaseTypeMysql:
		mysql.Init(&cfg.Mysql)
	}

	// 初始化redis
	redis.Init(&cfg.Redis)

	// 初始化Elasticsearch
	esearch.Init(&cfg.Elasticsearch)
}

func GetConfig() *models.Config {
	cfg := new(models.Config)
	configByte, err := os.ReadFile(models.FilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configByte, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
