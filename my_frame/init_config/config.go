package init_config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	DatabaseTypeMysql      = "mysql"
	DatabaseTypePostgresql = "postgresql"
)

type Config struct {
	Database struct {
		UseDatabase string         `yaml:"use-database"`
		Mysql       DatabaseConfig `yaml:"mysql"`
		Postgresql  DatabaseConfig `yaml:"postgresql"`
	} `yaml:"database"`
	Redis RedisConfig `yaml:"redis"`
}

type NamingStrategy struct {
	TableName          string
	SchemaName         string
	ColumnName         string
	JoinTableName      string
	RelationshipFKName string
	CheckerName        string
	IndexName          string
}

type DatabaseConfig struct {
	Host   string `yaml:"host"`    // 服务器地址
	DbName string `yaml:"db-name"` // 数据库名
	User   string `yaml:"user"`    // 用户名
	Pass   string `yaml:"pass"`    // 密码
	Port   int    `yaml:"port"`    // 端口
}

type RedisConfig struct {
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

const (
	FilePath = "./my_frame/config/config.yml"
)

var Cfg = &Config{}

func InitConfig() {
	configByte, err := os.ReadFile(FilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configByte, Cfg)
	if err != nil {
		panic(err)
	}
}
