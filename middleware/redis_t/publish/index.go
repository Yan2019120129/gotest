package publish

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/common/module/cache"
)

func main() {
	connect()
}

// connect 连接redis
func connect() {
	conn := cache.RdsPool.Get()
	defer conn.Close()
	var message string
	for {
		_, err := fmt.Scan(&message)
		if err != nil {
			return
		}
		publish(conn, message)
	}
}

// publish 发布信息
func publish(conn redis.Conn, message string) {
	// 发布消息到频道 "example_channel"
	reply, err := conn.Do("PUBLISH", "example_channel", message)
	if err != nil {
		fmt.Println("发布消息错误:", err)
		return
	}
	fmt.Printf("发布消息: %s\n", message)
	fmt.Printf("回复: %s\n", reply)
}
