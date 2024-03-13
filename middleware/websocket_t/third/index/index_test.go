package index

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
	"testing"
	"time"
)

const (
	IdEH         = "EH"
	ServerAddrEH = "wss://stream.talkfx.co/dconsumer/arrge"
	IdOne        = "7534801f-0919-4acb-abe8-3e46baa89ee9"
	IdTwo        = "244fbbc2-9055-4ce5-a6fd-65c69c88a054"
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
	//Instance.NewWs(uuidOkx, ServerOkxAddr).Run()
	//Instance.NewWs(IdEH, ServerAddrEH).Run(IdEH)

	fmt.Println(uuidValue, uuidTwo)
	time.Sleep(30 * time.Second)
}

// TestWebSocketOne 测试 websocket
func TestWebSocketOne(t *testing.T) {

	NewDefaultWs("wss://stream.talkfx.co/dconsumer/arrge").
		//SetSubMessage("connected").
		SetManage(&Manage{}).
		SetPulse(5).
		Run()

	//NewDefaultWs("wss://stream.talkfx.co/dconsumer/market/BTCUSD/1D").
	//	//SetSubMessage("connected").
	//	SetPulse(5).
	//	Run()
	time.Sleep(30 * time.Second)
	rds := cache.RdsPool.Get()
	defer rds.Close()
	// 查询热点数据
	hotData, err := redis.Strings(rds.Do("ZREVRANGE", "hot_data", 0, 10, "WITHSCORES"))
	if err != nil {
		fmt.Println("Error retrieving hot data:", err)
		return
	}

	// 输出热点数据
	for i := 0; i < len(hotData); i += 2 {
		member := hotData[i]
		score := hotData[i+1]
		fmt.Printf("Member: %s, Score: %s\n", member, score)
	}
}

type Manage struct {
}

func (m *Manage) DealWithMessage(msgType int, data []byte) {
	rds := cache.RdsPool.Get()
	defer rds.Close()
	tempData := make(map[string]interface{})
	err := json.Unmarshal(data, &tempData)
	if err != nil {
		logs.Logger.Info(logs.LogMsgApp, zap.Error(err))
		return
	}

	tableName := "hot_data"
	key := tempData["i"]
	if tempData["source"] == "1" {
		logs.Logger.Info(logs.LogMsgApp, zap.Reflect("data", tempData))
		_, err = rds.Do("ZINCRBY", tableName, 1, key)
		if err != nil {
			logs.Logger.Info(logs.LogMsgApp, zap.Error(err))
			return
		}
	}

}

//Member: EURSGD, Score: 68
//Member: GBPAUD, Score: 63
//Member: CHFJPY, Score: 57
//Member: AUDSGD, Score: 57
//Member: USDCNH, Score: 45
//Member: GBPNZD, Score: 43
//Member: USDDKK, Score: 42
//Member: GBPJPY, Score: 41
//Member: USDCHF, Score: 38
//Member: USDSGD, Score: 37
//Member: SGDJPY, Score: 37
