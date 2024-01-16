package redis_test_test

import (
	"gotest/middleware/redis_test"
	"sync"
	"testing"
	"time"
)

// Hset 插入哈希表
func TestHset(t *testing.T) {
	redis_test.Hset()
}

// TestHdel  删除哈希键值
func TestHdel(t *testing.T) {
	redis_test.Hdel()
}

// TestPublish 发布信息
func TestPublish(t *testing.T) {
	redis_test.Publish()
}

// TestSubscribe  订阅消息
func TestSubscribe(t *testing.T) {
	wg := sync.WaitGroup{}
	go redis_test.Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestSubscribe1  订阅消息
func TestSubscribe1(t *testing.T) {
	wg := sync.WaitGroup{}
	go redis_test.Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestGet 获取数据
func TestGetGet(t *testing.T) {
	for {
		redis_test.Get()
		time.Sleep(3 * time.Second)
	}
}
