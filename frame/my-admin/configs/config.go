package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	DatabaseTypeMysql      = "mysql"
	DatabaseTypePostgresql = "postgresql"
)

type Config struct {
	Gin           *GinConfig           `yaml:"gin"`
	GoAdmin       *GoAdminConfig       `yaml:"go-admin"`
	Redis         *RedisConfig         `yaml:"redis"`
	Elasticsearch *ElasticsearchConfig `yaml:"elasticsearch"`
}

// ElasticsearchConfig 配置参数。
type ElasticsearchConfig struct {
	IpAddress             []string `yaml:"ip-address"`              // IP地址
	MaxIdleConnsPerHost   int      `yaml:"max-idle-conns-per-host"` // 每个主机的最大空闲连接数
	ResponseHeaderTimeout int      `yaml:"response-header-timeout"` // 接收响应头的超时时间
	DialerTimeout         int      `yaml:"dialer-timeout"`          // 建立连接的上下文和超时设置
}

// GormConfig gorm 配置
type GormConfig struct {
	QueryFields   bool   `yaml:"query-fields"`   // 是否开启全字段匹配查询
	SingularTable bool   `yaml:"singular-table"` // 是否关闭数据库表复数s
	UseDatabase   string `yaml:"use-database"`   // 数据库选择
}

// GinConfig gin配置参数
type GinConfig struct {
	Port           string `yaml:"port"`             // gin端口号
	ReadTimeout    int    `yaml:"read-timeout"`     // 读取超时时间
	WriteTimeout   int    `yaml:"write-timeout"`    // 写入超时时间
	MaxHeaderBytes int    `yaml:"max-header-bytes"` // 请求头的最大字节数
}

// GoAdminConfig goAdmin 配置
type GoAdminConfig struct {
	ConfigPath string `yaml:"config-path"`
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
	Pool   RedisPoolConfig `yaml:"pool"`    // 连接池配置
}

// RedisPoolConfig redisPool 连接池连接配置
type RedisPoolConfig struct {
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

// 定义全局变量config，并初始化为nil
var cfg *Config

func init() {
	if cfg == nil {
		path := getConfigPath()
		configByte, err := os.ReadFile(path)
		if err != nil {
			log.Print("config init err:", err)
		}
		if err = yaml.Unmarshal(configByte, &cfg); err != nil {
			log.Print("config  unmarshal err：", err)
		}
	} else {
		log.Print("config  already exists！！！")
	}
}

// getConfigPath 获取配置文件路径
func getConfigPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Print("get config Caller path err！！！")
	}
	// 获取绝对路径
	absolutePath, err := filepath.Abs(file)
	if err != nil {
		log.Print("get config abs path err:", err)
	}

	// 修改为配置文件路径并返回
	return strings.ReplaceAll(absolutePath, ".go", ".yml")
}

// GetGin  获取gin 配置
func GetGin() *GinConfig {
	return cfg.Gin
}

// GetGoAdmin  获取goAdmin 配置
func GetGoAdmin() *GoAdminConfig {
	return cfg.GoAdmin
}

// GetRedis  获取redis 配置
func GetRedis() *RedisConfig {
	return cfg.Redis
}

// GetElasticsearch  获取elasticsearch 配置
func GetElasticsearch() *ElasticsearchConfig {
	return cfg.Elasticsearch
}
