package goredis

import (
	"sync"
	"testing"
	"time"
)

// Hset 插入哈希表
func TestHset(t *testing.T) {
	Hset()
}

// TestHdel  删除哈希键值
func TestHdel(t *testing.T) {
	Hdel()
}

// TestPublish 发布信息
func TestPublish(t *testing.T) {
	Publish()
}

// TestSubscribe  订阅消息
func TestSubscribe(t *testing.T) {
	wg := sync.WaitGroup{}
	go Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestSubscribe1  订阅消息
func TestSubscribe1(t *testing.T) {
	wg := sync.WaitGroup{}
	go Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestGet 获取数据
func TestGetGet(t *testing.T) {
	for {
		Get()
		time.Sleep(3 * time.Second)
	}
}
