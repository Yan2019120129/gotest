package goredis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gotest/common/module/logs"
)

// redis连接对象
var client *redis.Client

// ctx 创建上下文
var ctx = context.Background()

var MDTUSDT = map[string]string{
	"instType":  "SWAP",
	"instId":    "LTC-USD-SWAP",
	"last":      "9999.99",
	"lastSz":    "1",
	"askPx":     "9999.99",
	"askSz":     "11",
	"bidPx":     "8888.88",
	"bidSz":     "5",
	"open24h":   "9000",
	"high24h":   "10000",
	"low24h":    "8888.88",
	"volCcy24h": "2222",
	"vol24h":    "2222",
	"sodUtc0":   "0.1",
	"sodUtc8":   "0.1",
	"ts":        "1597026383085",
}

// init 初始化redis
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

// Hset 插入hash数据
func Hset() {

	data, err := json.Marshal(&MDTUSDT)
	if err != nil {
		panic(err)
	}
	key := "SPOT"
	if err = client.HSet(ctx, key, data).Err(); err != nil {
		panic(err)
	}

	if err = client.Close(); err != nil {
		panic(err)
	}
}

// Hdel 删除hash数据
func Hdel() {
	key := "SPOT"
	if err := client.Del(ctx, key).Err(); err != nil {
		panic(err)
	}
	spot := client.HGetAll(ctx, key)
	result, err := spot.Result()
	if err != nil {
		panic(err)
	}
	if err = client.Close(); err != nil {
		panic(err)
	}
	fmt.Println("result:", result)
}

// Publish 发布消息
func Publish() {
	channel := "SPOT"

	data, err := json.Marshal(&MDTUSDT)
	if err != nil {
		panic(err)
	}
	if err = client.Publish(ctx, channel, data).Err(); err != nil {
		panic(err)
	}
	for {
		var msg string
		if _, err = fmt.Scan(&msg); err != nil {
			panic(err)
		}
		if err = client.Publish(ctx, channel, msg).Err(); err != nil {
			panic(err)
		}
	}
}

// Subscribe 订阅消息
func Subscribe() {
	//channel := "SPOT"
	channel := "tickers-MDT-USDT"
	pubsub := client.Subscribe(ctx, channel)
	defer pubsub.Close()
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println("channel:", msg.Channel)
		fmt.Println("payload:", msg.Payload)
	}
}

// Get 获取缓存数据
func Get() {
	result := client.Get(ctx, "BTC-USDT")
	if result.Err() != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Err()))
	}
	logs.Logger.Info("信息", zap.Reflect("BTC-USDT", result.Val))

	result = client.Get(ctx, "ADA-USDT")
	if result.Err() != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Err()))
	}
	logs.Logger.Info("信息", zap.Reflect("ADA-USDT", result.Val))

	result = client.HGet(ctx, "tickers", "ADA-USDT")
	if result.Err() != nil {
		logs.Logger.Warn("错误信息", zap.Error(result.Err()))
	}
	logs.Logger.Info("信息", zap.Reflect("ADA-USDT", result.Val()))
}
