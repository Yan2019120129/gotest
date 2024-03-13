package subscribe

import (
	"testing"
	"time"
)

// TestSubscribe 测试redis订阅方法
func TestSubscribe(t *testing.T) {
	subscribe()
	time.Sleep(5 * time.Minute)
}
