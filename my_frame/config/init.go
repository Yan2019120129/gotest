package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"gotest/my_frame/models"
	"os"
	"sync"
)

// once 用于初始化config变量，并保证只初始化一次
var _once sync.Once

// 定义全局变量config，并初始化为nil
var config *models.Config

func init() {
	if config == nil {
		_once.Do(
			func() {
				if configByte, err := os.ReadFile(models.FilePath); err == nil {
					if err = yaml.Unmarshal(configByte, &config); err != nil {
						panic(err)
					}
					fmt.Printf("内存地址：%p----->配置文件初始化成功！！！\n", config)
				}
				panic(errors.New("配置文件初始化失败！！！"))
			},
		)
	} else {
		fmt.Println("配置文件实例已存在！！！")
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