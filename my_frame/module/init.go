package module

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gotest/my_frame/config"
	"os"
	"sync"
)

// once 用于初始化config变量，并保证只初始化一次
var _once sync.Once

// 定义全局变量config，并初始化为nil
var cfg *config.Config

func init() {
	if cfg == nil {
		_once.Do(
			func() {
				if configByte, err := os.ReadFile(config.FilePath); err == nil {
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
func GetGorm() *config.GormConfig {
	return &cfg.Gorm
}

// GetMysql  获取mysql 配置
func GetMysql() *config.DatabaseConfig {
	return &cfg.Gorm.Database.Mysql
}

// GetPostgres  获取postgres 配置
func GetPostgres() *config.DatabaseConfig {
	return &cfg.Gorm.Database.Postgresql
}

// GetGin  获取gin 配置
func GetGin() *config.GinConfig {
	return &cfg.Gin
}

// GetRedis  获取redis 配置
func GetRedis() *config.RedisConfig {
	return &cfg.Redis
}

// GetElasticsearch  获取elasticsearch 配置
func GetElasticsearch() *config.ElasticsearchConfig {
	return &cfg.Elasticsearch
}
