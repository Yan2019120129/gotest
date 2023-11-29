package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer c.Close()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	// 发送命令
	err = c.Send("HGET", "_Token", "1-0")
	if err != nil {
		fmt.Println("Error sending command:", err)
		return
	}

	// 接收结果
	result, err := redis.String(c.Receive())
	if err != nil {
		fmt.Println("Error receiving result:", err)
		return
	}

	fmt.Println("Result:", result)
}
