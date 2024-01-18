package okex

import (
	"go.uber.org/zap"
	"gotest/common/module/log/zap_log"
	"testing"
	"time"
)

// TestOkex
func TestOkex(t *testing.T) {
	// 订阅所有产品行情, 并且读取
	instance := NewOkexStruct()
	instance.SubscribeTickers().Reader()
	for {
		time.Sleep(3 * time.Second)
		BTC_USDT := instance.GetTicker("BTC-USDT")
		zap_log.Logger.Info("BTC_USDT", zap.Reflect("ticker", BTC_USDT))
		ADA_USDT := instance.GetTicker("ADA-USDT")
		zap_log.Logger.Info("ADA_USDT", zap.Reflect("ticker", ADA_USDT))
		//instance.TickerUpdatePrice()
		//fmt.Println("TickerUpdatePrice updated")
	}
}
