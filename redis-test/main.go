package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// redis连接对象
var client *redis.Client

func main() {
	TestHset()
}

// TestHset  插入哈希表
func TestHset() {
	MDTUSDT := map[string]string{
		"MDT-USDT": `{
			"instType":  "SPOT",
			"instId":    "MDT-USDT",
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
		}`,
	}
	ctx := context.Background()
	if err := client.HSet(ctx, "SPOT", MDTUSDT).Err(); err != nil {
		panic(err)
	}

	if err := client.HGetAll(ctx, "SPOT").Err(); err != nil {
		panic(err)
	}
}

// TestHdel  删除哈希表
func TestHdel() {
	ctx := context.Background()
	key := "SPOT"
	if err := client.Del(ctx, key).Err(); err != nil {
		panic(err)
	}
	spot := client.HGetAll(ctx, key)
	result, err := spot.Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("result:", result)
}

// init 初始化redis
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}
