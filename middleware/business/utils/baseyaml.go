package utils

import (
	"business/enum"
	"business/model"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func GetBaseConfig() (*model.Config, error) {
	baseFileByte, err := os.ReadFile(enum.PathBaseFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file:", err)
	}

	var config model.Config
	if err := yaml.Unmarshal(baseFileByte, &config); err != nil {
		return nil, fmt.Errorf("error parsing file: %v", err)
	}

	return &config, nil
}
