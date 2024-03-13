package redigo

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
)

// Get 获取缓存数据
func Get() {
	names := make([]string, 0)
	for i := 0; i < 10; i++ {
		names = append(names, gofakeit.LastName())
	}
	for i := 0; i < 10; i++ {
		go func() {
			rdsConn := cache.RdsPool.Get()
			defer rdsConn.Close()
			for {

				result, err := rdsConn.Do("HSET", "test", names[gofakeit.Number(0, 9)], gofakeit.Name())
				if err != nil {
					logs.Logger.Error("redis-go", zap.Error(err))
				}
				logs.Logger.Info("redis-go", zap.Reflect("result", result))
				//time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			rdsConn := cache.RdsPool.Get()
			defer rdsConn.Close()
			for {

				result, err := redis.Bytes(rdsConn.Do("HGET", "test", names[gofakeit.Number(0, 9)]))
				if err != nil {
					logs.Logger.Error("redis-go", zap.Error(err))
				}
				logs.Logger.Info("redis-go", zap.Reflect("result", string(result)))
				//time.Sleep(time.Millisecond * 500)
			}
		}()
	}
}
