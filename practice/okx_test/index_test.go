package okx_test_test

import (
	"go.uber.org/zap"
	"gotest/common/module/logger"
	"gotest/practice/okx_test"
	"testing"
	"time"
)

// TestGetTread 获取交易量
func TestGetTrades(t *testing.T) {
	for {
		data, err := okx_test.GetTrades("BTC-USDT")
		if err != nil {
			logger.Logger.Warn("warn", zap.Error(err))
		}
		logger.Logger.Info("info", zap.Reflect("data", data))
		time.Sleep(5 * time.Second)
	}
}
