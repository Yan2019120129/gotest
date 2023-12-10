package models

const (
	DatabaseTypeMysql      = "mysql"
	DatabaseTypePostgresql = "postgresql"
)

type Config struct {
	Gin           GinConfig     `yaml:"gin"`
	Database      Database      `yaml:"database"`
	Redis         RedisConfig   `yaml:"redis"`
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
}

// Elasticsearch 配置参数
type Elasticsearch struct {
	IpAddress             []string `yaml:"ip-address"`              // IP地址
	MaxIdleConnsPerHost   int      `yaml:"max-idle-conns-per-host"` // 每个主机的最大空闲连接数
	ResponseHeaderTimeout int      `yaml:"response-header-timeout"` // 接收响应头的超时时间
	DialerTimeout         int      `yaml:"dialer-timeout"`          // 建立连接的上下文和超时设置
}

type Database struct {
	UseDatabase string         `yaml:"use-database"`
	Mysql       DatabaseConfig `yaml:"mysql"`
	Postgresql  DatabaseConfig `yaml:"postgresql"`
}

// GinConfig gin配置参数
type GinConfig struct {
	Port           string `yaml:"port"`           // gin端口号
	ReadTimeout    int    `yaml:"readTimeout"`    // 读取超时时间
	WriteTimeout   int    `yaml:"writeTimeout"`   // 写入超时时间
	MaxHeaderBytes int    `yaml:"maxHeaderBytes"` // 请求头的最大字节数
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
	FilePath = "./my_frame/config.yml"
)
