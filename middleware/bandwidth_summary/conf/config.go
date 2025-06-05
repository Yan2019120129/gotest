package conf

import (
	"bandwidth_summary/utils"
	"fmt"
)

// 基础配置文件
type Config struct {
	Base   BaseConfig   `yaml:"base"`   // 基础配置
	Client ClientConfig `yaml:"client"` // 客户端配置
}

// 基础配置
type BaseConfig struct {
	TargetServer  string `yaml:"target_server"`  // 目标服务器
	CheckInterval int    `yaml:"check_interval"` // 检测间隔
}

type ClientConfig struct {
	Address []string `yaml:"address"` // 客户端地址
}

func LoadConf() (*Config, error) {
	var conf *Config
	isntance, err := utils.NewFileManager("./conf/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("Failed to load configuration file: " + err.Error())
	}

	// 转换为结构体
	err = isntance.ToStruct(conf)
	if err != nil {
		return nil, fmt.Errorf("Failed to struct err: " + err.Error())
	}

	// 检查必需的配置项
	if conf.Base.TargetServer == "" || conf.Base.CheckInterval <= 0 {
		panic("Invalid configuration: TargetServer and CheckInterval must be set")
	}

	return conf, nil
}
