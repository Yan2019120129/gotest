package config

import (
	"gopkg.in/yaml.v3"
	"gotest/my_frame/models"
	"os"
)

var config models.Config

func Init() {
	configByte, err := os.ReadFile(models.FilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configByte, &config)
	if err != nil {
		panic(err)
	}
}

// GetGorm  获取gorm 配置
func GetGorm() *models.GormConfig {
	return &config.Gorm
}

// GetMysql  获取mysql 配置
func GetMysql() *models.DatabaseConfig {
	return &config.Gorm.Database.Mysql
}

// GetPostgres  获取postgres 配置
func GetPostgres() *models.DatabaseConfig {
	return &config.Gorm.Database.Postgresql
}

// GetGin  获取gin 配置
func GetGin() *models.GinConfig {
	return &config.Gin
}

// GetRedis  获取redis 配置
func GetRedis() *models.RedisConfig {
	return &config.Redis
}

// GetElasticsearch  获取elasticsearch 配置
func GetElasticsearch() *models.ElasticsearchConfig {
	return &config.Elasticsearch
}
