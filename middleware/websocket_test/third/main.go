package main

import (
	"gotest/middleware/websocket_test/third/index"
	"sync"
)

const (
	// ServerOkxAddr okx 行情websocket 地址
	ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"

	// ServerAddrEH 易汇数据
	ServerAddrEH = "wss://stream.talkfx.co/dconsumer/arrge"

	// ServerCandleAndTradeAddr okx 行业websocket 地址
	ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"
)

var wg sync.WaitGroup

func main() {
	//data := []string{
	//	"{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"XRP-BTC\"\n    }]\n}",
	//	"{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"ETC-BTC\"\n    }]\n}",
	//}

	//for i, v := range data {
	index.NewDefaultWs(ServerAddrEH).
		//SetSubMessage("connected").
		SetPulse(5).
		Run()
	//index.NewDefaultWs(ServerOkxAddr).
	//	SetSubMessage(data...).
	//	SetPulse(5).
	//	Run()
	//}
	wg.Add(1)
	wg.Wait()
}
