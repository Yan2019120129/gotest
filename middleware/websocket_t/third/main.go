package main

import (
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"gotest/middleware/websocket_t/third/index"
	"strconv"
	"sync"
	"time"
)

const (
	// ServerOkxAddr okx 行情websocket 地址
	ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"

	// ServerAddrEH 易汇数据
	ServerAddrEH = "wss://stream.talkfx.co/dconsumer/arrge"

	// ServerAddrEBC 易汇数据
	ServerAddrEBC = "wss://stream.talkfx.co/dconsumer/market/BTCUSD/1D"

	// ServerCandleAndTradeAddr okx 行业websocket 地址
	ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"
)

var wg sync.WaitGroup

func main() {
	//index.NewDefaultWs(ServerAddrEH).
	//	//SetSubMessage("connected").
	//	SetPulse(5).
	//	Run()

	index.NewDefaultWs(ServerAddrEBC).
		SetSubMessage("{\n    \"cmd\": \"req\",\n    \"args\": [\n        \"candle.1D.BCHUSDT\",\n        361,\n        1678861513,\n        1709965573\n    ],\n    \"id\": \"trade.1D.BCHUSDT\"\n}").
		SetManage(&EH{}).
		SetPulse(0).
		Run()

	index.NewDefaultWs(ServerOkxAddr).
		SetSubMessage("{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"BCH-USDT\"\n    }]\n}").
		SetManage(&Okx{}).
		SetPulse(0).
		Run()

	wg.Add(1)
	wg.Wait()
}

type EH struct {
}

// DealWithMessage 处理易汇数据
func (e *EH) DealWithMessage(msgType int, data []byte) {
	ehData := EHData{}
	if err := json.Unmarshal(data, &ehData); err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	}

	if ehData.Data != nil {
		t := time.Unix(int64(ehData.Data.Dt)/1000, 0)
		date := t.Format("2006-01-02 15:04:05")
		logs.Logger.Info("EH", zap.String("date", date), zap.Float64("price", ehData.Data.Price))
	}
}

type EHData struct {
	Data *EHDetailsData `json:"data"`
	Type string         `json:"type"`
}

type EHDetailsData struct {
	Symbol string  `json:"symbol"`
	Prev   int     `json:"prev"`
	Type   string  `json:"type"`
	Cycle  string  `json:"cycle"`
	Dt     int     `json:"dt"`
	High   int     `json:"high"`
	Dhigh  float64 `json:"dhigh"`
	Low    int     `json:"low"`
	Price  float64 `json:"price"`
	Ask    float64 `json:"ask"`
	Dealer string  `json:"dealer"`
	Id     string  `json:"id"`
	Bid    float64 `json:"bid"`
	Close  int     `json:"close"`
	Open   float64 `json:"open"`
	Dlow   float64 `json:"dlow"`
}

type Okx struct {
}

// DealWithMessage 处理易汇数据
func (o *Okx) DealWithMessage(msgType int, msg []byte) {
	okxData := OkxData{}
	if err := json.Unmarshal(msg, &okxData); err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	}

	if okxData.Data != nil {
		data := okxData.Data[0]
		ts, err := strconv.ParseInt(data.Ts, 10, 64)
		if err != nil {
			return
		}
		t := time.Unix(ts/1000, 0)
		date := t.Format("2006-01-02 15:04:05")
		logs.Logger.Info("Okx", zap.String("date", date), zap.String("price", data.Last))
	}
}

type OkxData struct {
	Arg  *OkxDetailsArg    `json:"arg"`
	Data []*OkxDetailsData `json:"data"`
}

type OkxDetailsArg struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type OkxDetailsData struct {
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
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	Ts        string `json:"ts"`
}
