package main

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

// Hset 插入哈希表
func TestHset(t *testing.T) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		log.Println("redis dial failed.")
	}
	MDTUSDT := map[string]string{
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
	data, err := json.Marshal(&MDTUSDT)
	if err != nil {
		panic(err)
	}

	if err = conn.Send("hset", "SPOT", "MDT-USDT", data); err != nil {
		panic(err)
	}
}
