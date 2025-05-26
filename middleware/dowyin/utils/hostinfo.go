package utils

import (
	"encoding/json"
	"os"
)

func GetAppID() (string, error) {
	configFile := "/etc/xyapp/recruitResult.json"
	type ResultItem struct {
		AppID string `json:"appid"`
	}
	type Data struct {
		Result []ResultItem `json:"result"`
	}

	var ret string = "未知"
	if _, err := os.Stat(configFile); err == nil {
		data, err := os.ReadFile(configFile)
		if err == nil {
			var result Data
			if err := json.Unmarshal(data, &result); err == nil && len(result.Result) > 0 {
				ret = result.Result[0].AppID
			}
		} else {
			return ret, err
		}
	}
	return ret, nil
}
