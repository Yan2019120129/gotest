package config

const (
	Database_Type_Mysql      = "mysql"
	Database_Type_Postgresql = "postgresql"
)

const (
	FilePath = "/home/programmer-yan/Documents/GoFile/gotest/my_frame/config/config.yml"
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
	QueryFields   bool     `yaml:"query-fields"`   // 是否开启全字段匹配查询
	SingularTable bool     `yaml:"singular-table"` // 是否关闭数据库表复数s
	UseDatabase   string   `yaml:"use-database"`   // 数据库选择
	Database      Database `yaml:"database"`       // 数据库配置
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