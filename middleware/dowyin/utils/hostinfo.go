package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetAppID() (string, error) {
	configFile := "/etc/xyapp/recruitResult.json"
	//configFile := "/home/yan/Documents/file/gofile/gotest/common/file/recruitResult.json"
	type ResultItem struct {
		AppID string `json:"appid"`
	}
	type Data struct {
		Result []ResultItem `json:"result"`
	}

	if _, err := os.Stat(configFile); err == nil {
		data, err := os.ReadFile(configFile)
		if err != nil {
			return "", err
		}
		var result Data
		if err := json.Unmarshal(data, &result); err == nil && len(result.Result) > 0 {
			return result.Result[0].AppID, nil
		}
	}
	return "", fmt.Errorf("config file %s not found", configFile)
}
