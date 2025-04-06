package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"sync/atomic"
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

var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 原子操作版加函数
func atomicAdd() {
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&x, 1)
	}
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	for i := 0; i < 10; i++ {
		l.Lock()
		x += 1
		//x = gofakeit.Int64()
		l.Unlock()
	}
	wg.Done()
}

// 普通版加函数
func add() {
	for i := 0; i < 10; i++ {
		x += 1
		//x = gofakeit.Int64()
	}
	wg.Done()
}

// TestMutexAdd 测试原子操作添加数据
func TestMutexAdd(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add() // 普通版add函数 不是并发安全的
		//go mutexAdd() // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

// -2866055792857880426
// 20.549042ms

// -509795112266522614
//21.598417ms

// -537491006813644454
//17.001458ms
