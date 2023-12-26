package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/gen"
	"os"
	"sync"
)

const (
	DatabaseTypeMysql      = "mysql"
	DatabaseTypePostgresql = "postgresql"
)

var GenMdoe = map[string]interface{}{
	"WithoutContext":     gen.WithoutContext,
	"WithDefaultQuery":   gen.WithDefaultQuery,
	"WithQueryInterface": gen.WithQueryInterface,
}

const (
	//FilePath = "/home/programmer-yan/Documents/GoFile/gotest/my_frame/config/config.yml"
	FilePath = "/Users/taozi/Documents/Golang/gotest/my_frame/config/config.yml"
)

type Config struct {
	Gorm          GormConfig          `yaml:"gorm"`
	Gin           GinConfig           `yaml:"gin"`
	Redis         RedisConfig         `yaml:"redis"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}

// ElasticsearchConfig 配置参数。
type ElasticsearchConfig struct {
	IpAddress             []string `yaml:"ip-address"`              // IP地址
	MaxIdleConnsPerHost   int      `yaml:"max-idle-conns-per-host"` // 每个主机的最大空闲连接数
	ResponseHeaderTimeout int      `yaml:"response-header-timeout"` // 接收响应头的超时时间
	DialerTimeout         int      `yaml:"dialer-timeout"`          // 建立连接的上下文和超时设置
}

// Database 数据库配置
type Database struct {
	Mysql      DatabaseConfig `yaml:"mysql"`
	Postgresql DatabaseConfig `yaml:"postgresql"`
}

// GormConfig gorm 配置
type GormConfig struct {
	QueryFields   bool      `yaml:"query-fields"`   // 是否开启全字段匹配查询
	SingularTable bool      `yaml:"singular-table"` // 是否关闭数据库表复数s
	UseDatabase   string    `yaml:"use-database"`   // 数据库选择
	GenConfig     GenConfig `yaml:"gen"`            // 数据库配置
	Database      Database  `yaml:"database"`       // 数据库配置
}

// GenConfig gen配置
type GenConfig struct {
	OutPath string   `yaml:"out-path"` // 生成模型接口文件路径
	mode    []string `yaml:"mode"`     // 生成模式
}

// GinConfig gin配置参数
type GinConfig struct {
	Port           string `yaml:"port"`           // gin端口号
	ReadTimeout    int    `yaml:"readTimeout"`    // 读取超时时间
	WriteTimeout   int    `yaml:"writeTimeout"`   // 写入超时时间
	MaxHeaderBytes int    `yaml:"maxHeaderBytes"` // 请求头的最大字节数
}

// DatabaseConfig 数据库基本配置
type DatabaseConfig struct {
	Host   string `yaml:"host"`    // 服务器地址
	DbName string `yaml:"db-name"` // 数据库名
	User   string `yaml:"user"`    // 用户名
	Pass   string `yaml:"pass"`    // 密码
	Port   int    `yaml:"port"`    // 端口
}

// RedisConfig redis配置
type RedisConfig struct {
	UsePub bool            `yaml:"use-pub"` // 是否配置redis订阅功能
	Poll   RedisPollConfig `yaml:"poll"`    // 连接池配置
}

// RedisPollConfig redisPool 连接池连接配置
type RedisPollConfig struct {
	Network         string `yaml:"network"`            // 连接协议
	Server          string `yaml:"server"`             // 服务地址
	Port            int    `yaml:"port"`               // 端口号
	Pass            string `yaml:"pass"`               // 密码
	DbName          int    `yaml:"dbname"`             // 数据库
	ConnectTimeout  int    `yaml:"connect-timeout"`    // 连接超时时间
	ReadTimeout     int    `yaml:"read-timeout"`       // 读取超时时间
	WriteTimeout    int    `yaml:"Write-timeout"`      // 写入超时时间
	MaxOpenConn     int    `yaml:"max-open-conn"`      // 设置最大连接数
	ConnMaxIdleTime int    `yaml:"conn-max-idle-time"` // 空闲连接数
	MaxIdleConn     int    `yaml:"max-idle-conn"`      // 最大空闲连接数
	Wait            bool   `yaml:"wait"`               // 如果超时最大连接数是否等待
}

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

// GetRedisConfig  获取redis 配置
func GetRedisConfig() *RedisConfig {
	return &cfg.Redis
}

// GetElasticsearch  获取elasticsearch 配置
func GetElasticsearch() *ElasticsearchConfig {
	return &cfg.Elasticsearch
}
