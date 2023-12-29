package main

import (
	"sync"
	"testing"
)

// Hset 插入哈希表
func TestHset(t *testing.T) {
	Hset()
}

// TestHdel  删除哈希键值
func TestHdel(t *testing.T) {
	Hdel()
}

// TestSubscribe  订阅消息
func TestPublish(t *testing.T) {
	wg := sync.WaitGroup{}
	go Subscribe()
	wg.Add(1)
	wg.Wait()
}

// TestSubscribe1  订阅消息
func TestPublish1(t *testing.T) {
	wg := sync.WaitGroup{}
	go Subscribe()
	wg.Add(1)
	wg.Wait()
}
