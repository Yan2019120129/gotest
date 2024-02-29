package main

import (
	"github.com/google/uuid"
	"gotest/middleware/websocket_test/third/index"
	"sync"
)

const (
	// ServerOkxAddr okx 行情websocket 地址
	ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"

	// ServerCandleAndTradeAddr okx 行业websocket 地址
	ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"
)

var wg sync.WaitGroup

func main() {
	uuidValue := uuid.NewString()
	index.Instance.NewWs(uuidValue, ServerOkxAddr).
		SendMessage(uuidValue, []byte("{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"XRP-BTC\"\n    }]\n}")).
		Run()
	wg.Add(1)
	wg.Wait()
}
