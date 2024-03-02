package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"testing"
	"time"
)

// ConnText 上下文实例
type ConnText struct {
	ctx context.Context
	off context.CancelFunc
}

// TestGoroutine 测试协程的运作原理
func TestGoroutine(t *testing.T) {
	once := sync.Once{}
	ctxInstance, off := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		fmt.Println("测试")
		go Context(uuid.NewString(), ctxInstance)
		once.Do(func() {
			fmt.Println("我应该只运行一次")
		})
	}
	time.Sleep(5 * time.Second)
	off()
	time.Sleep(30 * time.Second)
}

// TestContext 测试Go上下文堵塞的问题
func Context(id string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(id)
		}
	}
}
