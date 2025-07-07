package conf

import (
	"fmt"
	"operate/utils"
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

var Conf Config

func InitConf(path string) error {
	isntance, err := utils.NewFileManager(path)
	if err != nil {
		return fmt.Errorf("failed to load configuration file: %s", err.Error())
	}

	// 转换为结构体
	err = isntance.YamlToStruct(&Conf)
	if err != nil {
		return fmt.Errorf("failed to struct err: %s", err.Error())
	}
	return nil
}
