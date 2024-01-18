package okx_test_test

import (
	"go.uber.org/zap"
	"gotest/common/module/log/zap_log"
	"gotest/practice/okx_test"
	"testing"
	"time"
)

// TestGetTread 获取交易量
func TestGetTrades(t *testing.T) {
	for {
		Books, err := okx_test.GetBooks("BTC-USDT")
		if err != nil {
			zap_log.Logger.Warn("warn", zap.Error(err))
		}
		zap_log.Logger.Info("info", zap.Reflect("Books", Books))

		kline, err := okx_test.GetKline("BTC-USDT", "3m")
		if err != nil {
			zap_log.Logger.Warn("warn", zap.Error(err))
		}
		zap_log.Logger.Info("info", zap.Reflect("kline", kline))

		Trades, err := okx_test.GetTrades("BTC-USDT")
		if err != nil {
			zap_log.Logger.Warn("warn", zap.Error(err))
		}
		zap_log.Logger.Info("info", zap.Reflect("Trades", Trades))
		time.Sleep(1 * time.Second)
	}
}
