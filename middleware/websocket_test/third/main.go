package main

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gotest/common/module/logs"
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
	data := []string{
		"{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"XRP-BTC\"\n    }]\n}",
		"{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"ETC-BTC\"\n    }]\n}",
	}

	for i, v := range data {
		uuidValue := uuid.NewString()
		logs.Logger.Info("run", zap.Int(uuidValue, i))
		index.Instance.SetWs(uuidValue, index.NewWs(ServerOkxAddr, 5, 5)).Run(uuidValue).SendMessage(uuidValue, []byte(v))
	}
	wg.Add(1)
	wg.Wait()
}
