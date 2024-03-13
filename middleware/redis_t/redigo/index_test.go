package redigo

import (
	"sync"
	"testing"
)

// TestRedisGoGet 测试Get 获取缓存数据方法
func TestRedisGoGet(t *testing.T) {
	wg := sync.WaitGroup{}
	Get()
	wg.Add(1)
	wg.Wait()
}
