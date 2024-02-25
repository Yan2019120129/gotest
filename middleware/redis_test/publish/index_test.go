package publish

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gotest/common/module/cache"
	"testing"
	"time"
)

// TestPublish 测试redis发布方法
func TestPublish(t *testing.T) {
	connect()
	time.Sleep(5 * time.Minute)
}

// TestPublish1 测试redis发布方法
func TestPublish1(t *testing.T) {
	connect()
	time.Sleep(5 * time.Minute)
}

// TestPublish2 测试redis发布方法
func TestPublish2(t *testing.T) {
	channel := []interface{}{"example_channel", "yan"}
	rdsConn := cache.RdsPubSubConn
	defer rdsConn.Close()
	dealWith := func(message []byte, channel string) {
		fmt.Println(channel + ":" + string(message))
	}

	go func() {
		if err := cache.Instance.Subscribe(dealWith, channel...); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		if err := cache.Instance.Subscribe(dealWith, "test"); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		cache.Instance.Publish([]byte(gofakeit.Name()), channel...)
		cache.Instance.Publish([]byte(gofakeit.Name()), "test")
		time.Sleep(3 * time.Second)
	}
}
