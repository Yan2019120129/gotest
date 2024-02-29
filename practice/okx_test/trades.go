package okx

import (
	"basic/module/cache"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gomodule/redigo/redis"
	"net/url"
)

// TradesData 交易数据
type TradesData struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Ts      string `json:"ts"`
	Count   string `json:"count"`
}

// trades 返回的推送数据
type trades struct {
	Arg  Arg          `json:"Arg"`
	Data []TradesData `json:"data"`
}

// GetTrades 交易量
func GetTrades(instId string) ([]*TradesData, error) {
	// 发送请求获取交易深度数据
	query := url.Values{"instId": {instId}}
	resp, err := Get(serverOkxAddrMap[serverAddrTrades], query)
	if err != nil {
		return nil, err
	}

	data := make([]*TradesData, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}

// GetWsTrades 获取交易深度数据
func GetWsTrades(instId string) (*TradesData, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	result, err := redis.Bytes(rdsConn.Do("HGET", ChannelMap[ChannelTrades], instId))
	if err != nil {
		log.Errorw("okx", "method", "GetRdsTicker", "error", err)
		return nil, err
	}

	// 解析redis 数据
	data := &TradesData{}
	_ = json.Unmarshal(result, &data)

	return data, nil
}
