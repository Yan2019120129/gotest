package okx

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
	"gotest/common/utils"
	"strconv"
)

// ticker 返回的推送数据
type ticker struct {
	Arg  Arg          `json:"Arg"`
	Data []tickerData `json:"data"`
}

// tickerData 行情推送的数据
type tickerData struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        string `json:"ts"`
}

// TickerData 用于存储到用户外部方便使用
type TickerData struct {
	InstId    string  `json:"instId"`
	Last      float64 `json:"last"`
	LastSz    float64 `json:"lastSz"`
	Open24h   float64 `json:"open24h"`
	High24h   float64 `json:"high24h"`
	Low24h    float64 `json:"low24h"`
	VolCcy24h float64 `json:"volCcy24h"` // 24小时成交量，以币为单位。 如果是衍生品合约，数值为交易货币的数量。 如果是币币/币币杠杆，数值为计价货币的数量。
	Vol24h    float64 `json:"vol24h"`    // 24小时成交量，以张为单位,如果是衍生品合约，数值为合约的张数。如果是币币/币币杠杆，数值为交易货币的数量。
	Ts        int     `json:"ts"`
}

// GetWsTicker 获取行情数据
func GetWsTicker(instId string) (*TickerData, error) {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()
	result, err := redis.Bytes(rdsConn.Do("HGET", ChannelMap[ChannelTicker], instId))
	if err != nil {
		logs.Logger.Error(logs.LogMsgOkx, zap.Error(err))
		return nil, err
	}

	// 解析redis 数据
	data := tickerData{}
	if err = json.Unmarshal(result, &data); err != nil {
		logs.Logger.Error("okx", zap.Error(err))
	}

	// 转换为对应的类型
	tickerConvertData := &TickerData{}
	tickerConvertData.InstId = data.InstId
	tickerConvertData.Last, _ = strconv.ParseFloat(data.Last, 64)
	tickerConvertData.LastSz, _ = strconv.ParseFloat(data.LastSz, 64)
	tickerConvertData.Open24h, _ = strconv.ParseFloat(data.Open24h, 64)
	tickerConvertData.High24h, _ = strconv.ParseFloat(data.High24h, 64)
	tickerConvertData.Low24h, _ = strconv.ParseFloat(data.Low24h, 64)
	tickerConvertData.VolCcy24h, _ = strconv.ParseFloat(data.VolCcy24h, 64)
	tickerConvertData.Vol24h, _ = strconv.ParseFloat(data.Vol24h, 64)
	tickerConvertData.Ts, _ = strconv.Atoi(data.Ts)

	return tickerConvertData, nil
}

// GetTickerString 获取行情数据
func GetTickerString(instId string) (string, error) {
	data, err := GetWsTicker(instId)
	if err != nil {
		return "", err
	}

	dataString := utils.ObjToString(data)
	return dataString, nil
}
