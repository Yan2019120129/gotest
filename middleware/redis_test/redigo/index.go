package redigo

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
	"time"
)

// Get 获取缓存数据
func Get() {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()
	for {
		result, err := redis.Bytes(rdsConn.Do("HGET", "tickers", "ADA-USDT"))
		if err != nil {
			logs.Logger.Error("redis-go", zap.Error(err))
		}
		logs.Logger.Info("redis-go", zap.Reflect("result", string(result)))
		time.Sleep(time.Second * 3)
	}
}
