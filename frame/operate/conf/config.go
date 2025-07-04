package conf

import (
	"fmt"
	"tx_script/utils"
)

// 基础配置文件
type Config struct {
	Base BaseConfig `yaml:"base"` // 基础配置
	Log  Log        `yaml:"log"`  // 日志基本配置
}

// 基础配置
type BaseConfig struct {
	TargetServer  string `yaml:"target_server"`  // 目标服务器
	CheckInterval int    `yaml:"check_interval"` // 检测间隔
	SDK           string `yaml:"sdk"`            // 机房SDK
}

// 日志配置
type Log struct {
	Dir        string `yaml:"dir"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

func LoadConf(path string) (*Config, error) {
	conf := Config{}
	isntance, err := utils.NewFileManager(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration file: %s", err.Error())
	}

	// 转换为结构体
	err = isntance.YamlToStruct(&conf)
	if err != nil {
		return nil, fmt.Errorf("failed to struct err: %s", err.Error())
	}

	// 检查必需的配置项
	if conf.Base.TargetServer == "" || conf.Base.SDK == "" {
		return nil, fmt.Errorf("invalid configuration: TargetServer and SDK must be set")
	}

	return &conf, nil
}
