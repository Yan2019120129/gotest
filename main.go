package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

func NewRedisClient() (conn redis.Conn, err error) {
	host := "127.0.0.1"
	port := "6379"
	adderss := host + ":" + port
	c, err := redis.Dial("tcp", adderss)
	return c, err
}

func ResolveOrderCreate(wait *sync.WaitGroup) {
	defer wait.Done()
	conn, err := NewRedisClient()
	if err != nil {
		return
	}
	client := redis.PubSubConn{conn}
	err = client.Subscribe("order-create")
	if err != nil {
		fmt.Println("订阅错误:", err)
		return
	}
	fmt.Println("等待订阅数据 ---->")
	for {
		switch v := client.Receive().(type) {
		case redis.Message:
			fmt.Println("Message", v.Channel, string(v.Data))
		case redis.Subscription:
			fmt.Println("Subscription", v.Channel, v.Kind, v.Count)
		}
	}
}

func Publish() {
	conn, err := NewRedisClient()
	if err != nil {
		return
	}
	type Data struct {
		Name *string
		Age  *int
	}
	data := &Data{}
	name := "波兰中锋 周琦"
	age := 25
	data.Name = &name
	data.Age = &age
	_, err = conn.Do("Publish", "order-create", "1111111111111")
	if err != nil {
		fmt.Println("发布错误", err)
		return
	}
	_, err = conn.Do("Publish", "order-create", 123)
	if err != nil {
		fmt.Println("发布错误", err)
		return
	}
	_, err = conn.Do("Publish", "order-create", data)
	if err != nil {
		fmt.Println("发布错误", err)
		return
	}
	_, err = conn.Do("Publish", "order-create", "33333333333333")
	if err != nil {
		fmt.Println("发布错误", err)
		return
	}
	_, err = conn.Do("Publish", "order-create", "66666666666666")
	if err != nil {
		fmt.Println("发布错误", err)
		return
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go ResolveOrderCreate(&wg)
	Publish()
	wg.Wait()
}
