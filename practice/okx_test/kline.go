package okx

import (
	"encoding/json"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"net/url"
)

// KlineBarList 时间粒度，默认值1m
var KlineBarList = []*KlineBar{
	{Label: "1M", Value: "1m"},
	{Label: "5M", Value: "5m"},
	{Label: "30M", Value: "30m"},
	{Label: "1H", Value: "1H"},
	{Label: "4H", Value: "4H"},
	{Label: "1D", Value: "1D"},
}

type KlineBar struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// KlineData k线图数据
type KlineData struct {
	OpenPrice  float64 `json:"openPrice"`  //开盘价格
	HighPrice  float64 `json:"highPrice"`  //最高价格
	LowsPrice  float64 `json:"lowsPrice"`  //最低价格
	ClosePrice float64 `json:"closePrice"` //收盘价格
	Vol        float64 `json:"vol"`        //交易量
	Amount     float64 `json:"amount"`     //成交额
	CreatedAt  int64   `json:"createdAt"`  //开盘时间
}

// GetKline 获取k线图数据
func GetKline(instId, bar string) ([]*KlineData, error) {
	query := url.Values{"instId": {instId}, "limit": {"300"}, "bar": {bar}}
	resp, err := Get(serverOkxAddrMap[serverAddrCandles], query)
	if err != nil {
		return nil, err
	}

	data := make([]*KlineData, 0)
	err = json.Unmarshal(resp, &data)
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	}
	return data, nil
}
