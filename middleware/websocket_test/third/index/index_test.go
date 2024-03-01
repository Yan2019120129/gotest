package index

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

const (
	// ServerOkxAddr okx 行情websocket 地址
	ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"

	// ServerCandleAndTradeAddr okx 行业websocket 地址
	ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"
)

// TestWebSocket 测试websocket
func TestWebSocket(t *testing.T) {
	uuidValue := uuid.NewString()
	Instance.NewWs(uuidValue, ServerOkxAddr).Run().SendMessage(uuidValue, []byte("{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"XRP-BTC\"\n    }]\n}"))
	time.Sleep(30 * time.Second)
}
