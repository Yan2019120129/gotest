package dto

import "encoding/json"

// RespJson 返回错误
type RespJson struct {
	Code string          `json:"code"` //	状态码
	Msg  string          `json:"msg"`  //	错误消息
	Data json.RawMessage `json:"data"` //	数据
}

// SubscribeRespJson 订阅返回数据
type SubscribeRespJson struct {
	Arg  *SubscribeArg   `json:"arg"`
	Data json.RawMessage `json:"data"`
}

// SubscribeArg 币种订阅频道。
type SubscribeArg struct {
	Channel string `json:"channel"` // 订阅的通道
	InstID  string `json:"instId"`  // 货币类型
}

// OkxTickers 产品行情信息
type OkxTickers struct {
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
