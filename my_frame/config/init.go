package config

import (
	"gopkg.in/yaml.v3"
	"gotest/my_frame/models"
	"os"
)

func GetConfig() *models.Config {
	cfg := new(models.Config)
	configByte, err := os.ReadFile(models.FilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configByte, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
