package okxserver

import (
	"github.com/gorilla/websocket"
	"gotest/frame/my_frame/module/logs"
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
		logs.Logger.Error(err.Error())
	}
	defer conn.Close()

	//rdsConn := cache.RdsPubSubConn
	//defer rdsConn.Close()
	//ctx := context.Background()
	//if err := rdsConn.Subscribe(ctx, "tickers-MDT-USDT"); err != nil {
	//	logger.Logger.Error(err.Error())
	//}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// 向客户端发送消息
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			logs.Logger.Error(err.Error())
		}

		//switch v := rdsConn.Receive().(type) {
		//case redis.Message:
		//	logger.Logger.Info(string(v.Data))
		//
		//
		//case redis.Subscription:
		//	logger.Logger.Info(v.Channel + ":" + strconv.Itoa(v.Count))
		//case error:
		//	logger.Logger.Error(v.Error())
		//	return
		//}
	}
}
