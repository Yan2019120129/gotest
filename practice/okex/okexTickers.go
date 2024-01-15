package okex

import (
	"encoding/json"
	"strconv"
)

const OkexTickersURL = ""

type Ticker struct {
	InstType  string `json:"instType"`  //	产品类型
	InstId    string `json:"instId"`    //	产品ID
	Last      string `json:"last"`      //	最新价
	LastSz    string `json:"lastSz"`    //	最新成交的数量
	AskPx     string `json:"askPx"`     //	卖一价
	AskSz     string `json:""`          //	卖一价挂单数量
	BidPx     string `json:"bidPx"`     //	买一价
	BidSz     string `json:"bidSz"`     //	买一价挂单数量
	Open24h   string `json:"open24h"`   // 	24小时开盘价
	High24h   string `json:"high24h"`   //	24小时最高价
	Low24h    string `json:"low24h"`    //	24小时最低价
	VolCcy24h string `json:"volCcy24h"` //	24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    //	24小时成交量，以张为单位
	SodUtc0   string `json:"sodUtc0"`   //	UTC 0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   //	UTC+8 时开盘价
	Ts        string `json:"ts"`        //	ticker数据产生时间
}

// GetLast 获取最新价格
func (_Ticker *Ticker) GetLast() float64 {
	lastFloat64, _ := strconv.ParseFloat(_Ticker.Last, 64)
	return lastFloat64
}

// GetLastSz 获取最新成交量
func (_Ticker *Ticker) GetLastSz() float64 {
	lastSzFloat64, _ := strconv.ParseFloat(_Ticker.LastSz, 64)
	return lastSzFloat64
}

// GetOpen24h 获取24小时开盘价
func (_Ticker *Ticker) GetOpen24h() float64 {
	open24hFloat64, _ := strconv.ParseFloat(_Ticker.Open24h, 64)
	return open24hFloat64
}

// GetHigh24h 获取24小时最高价
func (_Ticker *Ticker) GetHigh24h() float64 {
	high24hFloat64, _ := strconv.ParseFloat(_Ticker.High24h, 64)
	return high24hFloat64
}

// GetLow24h 获取24小时最低价
func (_Ticker *Ticker) GetLow24h() float64 {
	lows24hFloat64, _ := strconv.ParseFloat(_Ticker.Low24h, 64)
	return lows24hFloat64
}

// GetVol24h 获取24小时成交量
func (_Ticker *Ticker) GetVol24h() float64 {
	vol24hFloat64, _ := strconv.ParseFloat(_Ticker.Vol24h, 64)
	return vol24hFloat64
}

// GetAmount24h 获取24小时成交额
func (_Ticker *Ticker) GetAmount24h() float64 {
	amount24hFloat64, _ := strconv.ParseFloat(_Ticker.VolCcy24h, 64)
	return amount24hFloat64
}

// GetTs 获取当前时间戳
func (_Ticker *Ticker) GetTs() int64 {
	ts, _ := strconv.ParseInt(_Ticker.Ts, 10, 64)
	return ts
}

// Tickers 获取所有行情
func (_OkexStruct *OkexStruct) Tickers() ([]*Ticker, error) {
	params := map[string]interface{}{"instType": "SPOT"}
	resp, err := _OkexStruct.Get("/api/v5/market/tickers", params)
	if err != nil {
		return nil, err
	}

	data := make([]*Ticker, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}

// Ticker 获取单币行情
func (_OkexStruct *OkexStruct) Ticker(symbol string) ([]*Ticker, error) {
	params := map[string]interface{}{"instId": symbol}
	resp, err := _OkexStruct.Get("/api/v5/market/ticker", params)
	if err != nil {
		return nil, err
	}

	data := make([]*Ticker, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}
