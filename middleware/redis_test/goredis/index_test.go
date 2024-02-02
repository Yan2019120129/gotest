package goredis_test

import (
	"gotest/middleware/redis_test/goredis"
	"sync"
	"testing"
	"time"
)

// Hset 插入哈希表
func TestHset(t *testing.T) {
	goredis.Hset()
}

// TestHdel  删除哈希键值
func TestHdel(t *testing.T) {
	goredis.Hdel()
}

// TestPublish 发布信息
func TestPublish(t *testing.T) {
	goredis.Publish()
}

// TestSubscribe  订阅消息
func TestSubscribe(t *testing.T) {
	wg := sync.WaitGroup{}
	go goredis.Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestSubscribe1  订阅消息
func TestSubscribe1(t *testing.T) {
	wg := sync.WaitGroup{}
	go goredis.Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestGet 获取数据
func TestGetGet(t *testing.T) {
	for {
		goredis.Get()
		time.Sleep(3 * time.Second)
	}
}
