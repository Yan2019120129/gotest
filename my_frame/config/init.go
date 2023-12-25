package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

// once 用于初始化config变量，并保证只初始化一次
var _once sync.Once

// 定义全局变量config，并初始化为nil
var cfg *Config

func init() {
	if cfg == nil {
		_once.Do(
			func() {
				if configByte, err := os.ReadFile(FilePath); err == nil {
					if err = yaml.Unmarshal(configByte, &cfg); err != nil {
						panic(err)
					}
					fmt.Printf("内存地址：%p----->配置文件初始化成功！！！\n", cfg)
				} else {
					panic(err)
				}
			},
		)
	} else {
		fmt.Println("配置文件实例已存在！！！")
	}
}

// GetGorm  获取gorm 配置
func GetGorm() *GormConfig {
	return &cfg.Gorm
}

// GetGen  获取gen 配置
func GetGen() *GenConfig {
	return &cfg.Gorm.GenConfig
}

// GetMysql  获取mysql 配置
func GetMysql() *DatabaseConfig {
	return &cfg.Gorm.Database.Mysql
}

// GetPostgres  获取postgres 配置
func GetPostgres() *DatabaseConfig {
	return &cfg.Gorm.Database.Postgresql
}

// GetGin  获取gin 配置
func GetGin() *GinConfig {
	return &cfg.Gin
}

// GetRedis  获取redis 配置
func GetRedis() *RedisConfig {
	return &cfg.Redis
}

// GetElasticsearch  获取elasticsearch 配置
func GetElasticsearch() *ElasticsearchConfig {
	return &cfg.Elasticsearch
}
