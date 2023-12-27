package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// sub 订阅通道数据
func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		log.Println("redis dial failed.")
	}

	// 创建订阅者连接
	psc := redis.PubSubConn{Conn: conn}

	// 订阅频道 "example_channel"
	if err := psc.Subscribe("example_channel"); err != nil {
		fmt.Println("订阅频道错误:", err)
		return
	}

	fmt.Println("等待接收消息...")

	// 循环接收消息
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("收到消息: %s\n", v.Data)
		case redis.Subscription:
			fmt.Printf("订阅频道: %s，订阅数量: %d\n", v.Channel, v.Count)
		case error:
			fmt.Println("接收消息错误:", v)
			return
		}
	}
}
