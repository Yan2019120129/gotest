package conf

import (
	"embed"
	"gopkg.in/yaml.v3"
)

// 基础配置文件
type Config struct {
	Base BaseConfig `yaml:"base"` // 基础配置
	Log  Log        `yaml:"log"`  // 日志基本配置
}

// 基础配置
type BaseConfig struct {
	TargetServer string     `yaml:"target_server"` // 目标服务器
	JumpServer   JumpServer `yaml:"jump_server"`   // jumpserver用户名
}

// JumpServer 配置
type JumpServer struct {
	Port string `yaml:"port"`
	User string `yaml:"user"`
}

// 日志配置
type Log struct {
	Dir        string `yaml:"dir"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

var Conf Config

func InitConf(fs embed.FS) error {
	file, err := fs.ReadFile("conf/config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		return err
	}

	return nil
}
