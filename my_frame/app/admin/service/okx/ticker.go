package okxserver

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"gotest/my_frame/module/cache"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 在这里进行一些逻辑判断，例如检查请求的域名是否在白名单中
		return true // 或者根据判断返回相应的结果
	},
}

// TickerIndex 获取行情数据
func TickerIndex(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket 升级错误:", err)
		panic(err)
		return
	}
	defer conn.Close()

	rdsConn := cache.RdsPubSubConn
	defer rdsConn.Close()
	ctx := context.Background()
	if err := rdsConn.Subscribe(ctx, "tickers-MDT-USDT"); err != nil {
		panic(err)
		return
	}

	for {
		switch v := rdsConn.Receive().(type) {
		case redis.Message:
			fmt.Printf("收到消息: %s\n", v.Data)
			// 向客户端发送消息
			if err := conn.WriteMessage(websocket.TextMessage, v.Data); err != nil {
				fmt.Println("服务端发送消息错误:", err)
				panic(err)
				return
			}
		case redis.Subscription:
			fmt.Printf("订阅频道: %s，订阅数量: %d\n", v.Channel, v.Count)
		case error:
			fmt.Println("接收消息错误:", v)
			panic(v)
			return
		}

	}
}
