package index

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

const (
	IdOne = "7534801f-0919-4acb-abe8-3e46baa89ee9"
	IdTwo = "244fbbc2-9055-4ce5-a6fd-65c69c88a054"
)

// TestWebSocket 测试websocket
func TestWebSocket(t *testing.T) {
	// 为什么这一部分没有执行
	uuidValue := uuid.NewString()
	uuidTwo := uuid.NewString()
	//Instance.NewWs(uuidValue, ServerOkxAddr).
	//	Run()
	//Instance.NewWs(uuidValue, ServerOkxAddr).
	//	Run().
	//	SendMessage(&Massage{
	//		Id:   uuidValue,
	//		Type: WsMessageTypeSub,
	//		Data: []byte("{\n    \"op\": \"subscribe\",\n    \"args\": [{\n        \"channel\": \"tickers\",\n        \"instId\": \"ETC-BTC\"\n    }]\n}"),
	//	})
	Instance.NewWs(uuidOkx, ServerOkxAddr).Run()
	fmt.Println(uuidValue, uuidTwo)
	time.Sleep(30 * time.Second)
}
