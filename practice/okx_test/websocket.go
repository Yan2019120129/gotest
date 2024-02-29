package okx

import (
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/module/logger"
	"basic/module/socket"
	"basic/utils"
	"github.com/fasthttp/websocket"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"strings"
	"time"
)

// WsInstance websocket实例。
type WsInstance struct {
	Conn             *websocket.Conn // websocket 实例
	MaxReconnect     int             // 设置最大重连次数
	ServerAddr       string          // 连接地址
	SubscribeChannel []Channel       // 订阅通道
	isClose          chan bool       // 用于关闭协程
	data             chan string     // 用于接收，处理数据。
}

// OkxParams 发送参数。
type OkxParams struct {
	Op   string `json:"op"`   // 操作，subscribe unsubscribe
	Args []*Arg `json:"args"` // 请求订阅的频道列表
}

// Arg 币种订阅频道。
type Arg struct {
	Channel Channel `json:"channel"` // 订阅的通道
	InstID  string  `json:"instId"`  // 货币类型
}

// SubscribeData 订阅Okx数据
type SubscribeData struct {
	Arg  Arg           `json:"Arg"`
	Data []interface{} `json:"data"`
}

// Data 返回给客户端的数据
type Data struct {
	Arg
	Data interface{} `json:"data"`
}

// NewOkx 创建okx 实例
func NewOkx() *WsInstance {
	instance := &WsInstance{
		Conn:             new(websocket.Conn), // websocket 实例
		MaxReconnect:     5,                   // 最大重连次数
		ServerAddr:       ServerOkxAddr,       // 订阅的通道，用于使用指定的连接地址
		SubscribeChannel: make([]Channel, 0),
		isClose:          make(chan bool),        // 用于关闭协程
		data:             make(chan string, 100), // 数据传输通道，用户接收处理数据
	}
	return instance
}

// Run 启动websocket
func (ws *WsInstance) Run() {
	// 链接websocket
	if err := ws.connect(); err != nil {
		logger.Logger.Info("run close")
		logger.Logger.Warn(logger.LogMsgOkx, zap.Error(err))
		return
	}

	// 发送测试数据
	ws.sendMessage()

	// 心跳检测
	go ws.heartbeat()

	// 读取信息
	go ws.read()

	// 订阅信息
	go ws.subscribe()

	// 处理并发送信息
	go ws.publish()

	logger.Logger.Info("okx run")
}

// Close 关闭链接
func (ws *WsInstance) Close() {
	defer logger.Logger.Info("okx close")
	//ws.closeMsg <- "close"
	//close(ws.isClose)
	close(ws.data)
	// 关闭链接
	if err := ws.Conn.Close(); err != nil {
		logger.Logger.Warn("Close", zap.Error(err))
		return
	}
}

// ConnectWS 连接okx websocket。
func (ws *WsInstance) connect() (err error) {
	// 添加协程使用WaitGroup管理线程状态
	ws.Conn, _, err = websocket.DefaultDialer.Dial(ws.ServerAddr, nil)
	if err != nil {
		logger.Logger.Warn(logger.LogMsgOkx, zap.Error(err))
		return err
	}

	logger.Logger.Info(logger.LogMsgOkx, zap.String("connect", "Connection completed"))
	return nil
}

// SendMessage 发送消息。
func (ws *WsInstance) SendMessage(message []byte) {
	if err := ws.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		logger.Logger.Warn(logger.LogMsgOkx, zap.Error(err))
	}
}

// read 读取消息。
func (ws *WsInstance) read() {
	logger.Logger.Info("read run")
	defer logger.Logger.Info("read close")
	for {
		select {
		// 当关闭通道的时候退出协程
		case <-ws.isClose:
			return
		default:
			// 当读取消息发送错误，则重新连接
			_, message, err := ws.Conn.ReadMessage()
			if err != nil {
				logger.Logger.Error(logger.LogMsgOkx, zap.Error(err))
				ws.reconnect()
				continue
			}
			ws.data <- string(message)
		}
	}
}

// publish 处理发布信息。
func (ws *WsInstance) publish() {
	logger.Logger.Info("publish run")
	defer logger.Logger.Info("publish close")
	rds := cache.Rds.Get()
	defer func(rds redis.Conn) {
		if err := rds.Close(); err != nil {
			return
		}
	}(rds)
	for {
		select {
		// 当关闭通道的时候退出协程
		case <-ws.isClose:
			return
		default:
			message := <-ws.data
			if strings.Contains(message, "error") && strings.Contains(message, "msg") {
				logger.Logger.Error(logger.LogMsgOkx, zap.String("message", message))
				continue
			}

			subscribeData := SubscribeData{}
			if err := json.Unmarshal([]byte(message), &subscribeData); err == nil {
				if subscribeData.Data != nil {
					cache.Instance.Publish(string(subscribeData.Arg.Channel), utils.JsonToBytes(&Data{
						Arg: Arg{
							Channel: subscribeData.Arg.Channel,
							InstID:  subscribeData.Arg.InstID,
						},
						Data: subscribeData.Data[0],
					}))
					//logger.Logger.Info(logger.LogMsgOkx, zap.Any("channel", subscribeData.Arg.Channel), zap.String("instType", subscribeData.Arg.InstID), zap.Any("data", subscribeData.Data[0]))
					_, err = rds.Do("HSET", subscribeData.Arg.Channel, subscribeData.Arg.InstID, utils.JsonToBytes(subscribeData.Data[0]))
					if err != nil {
						logger.Logger.Error(logger.LogMsgOkx, zap.Error(err))
					}
				}
				continue
			}
			logger.Logger.Warn(logger.LogMsgOkx, zap.String("message", message))
		}
	}
}

// heartbeat 每个几秒尝试发送连接消息，检测连接心跳
func (ws *WsInstance) heartbeat() {
	logger.Logger.Info("heartbeat run")
	defer logger.Logger.Info("heartbeat close")
	for {
		select {
		// 当关闭通道的时候退出协程
		case <-ws.isClose:
			return
		default:

			// 发送心跳消息
			ws.SendMessage([]byte("ping"))

			// 每隔三秒发送消息
			time.Sleep(5 * time.Second)
		}
	}
}

// reconnect 重新连接
func (ws *WsInstance) reconnect() {
	// 根觉最大的连接次数，断开连接的时候重新连接
	defer logger.Logger.Info("reconnect close")
	for i := 0; i < ws.MaxReconnect; i++ {
		logger.Logger.Warn(logger.LogMsgOkx, zap.Int("reconnectNum", ws.MaxReconnect))

		if err := ws.connect(); err != nil {
			// 进行MaxReconnect次连接之后仍然不成功则退出所有协程
			ws.MaxReconnect--
			logger.Logger.Error(logger.LogMsgOkx, zap.Int("reconnectNum", ws.MaxReconnect), zap.Error(err))
			continue
		}

		ws.sendMessage()
		return
	}
	ws.Close()
}

// subscribe 订阅消息。
func (ws *WsInstance) subscribe() {
	defer logger.Logger.Info("subscribe close")
	for {
		select {
		// 当关闭通道的时候退出协程
		case <-ws.isClose:
			return
		default:
			for _, channel := range ws.SubscribeChannel {
				ws.SetSubscribe(channel)
			}
		}
	}
}

// SetSubscribe 设置订阅
func (ws *WsInstance) SetSubscribe(SubscribeChannel Channel) {
	// 启动订阅模式
	channelMessageFunc := func(data []byte) {
		// 获取订阅该频道的uuid
		uuidList := socket.Instance.GetSubscribes(string(SubscribeChannel))

		msg := Data{}
		if err := json.Unmarshal(data, &msg); err != nil {
			logger.Logger.Error(logger.LogMsgOkx, zap.Error(err))
			return
		}
		for uuid, subscribe := range uuidList {
			for _, instType := range subscribe.Args {
				if msg.Arg.Channel == SubscribeChannel && instType == msg.Arg.InstID {
					logger.Logger.Info(logger.LogMsgOkx, zap.String("UUID", uuid), zap.Any("channel", msg.Arg.Channel), zap.String("instType", msg.Arg.InstID), zap.Reflect("data", msg.Data))
					_ = socket.Instance.WriteJSON(uuid, msg)
				}
			}
		}
	}

	_ = cache.Instance.Subscribe(string(SubscribeChannel), channelMessageFunc)
}

// sendTickerAll 发送订阅okx的信息
func (ws *WsInstance) sendMessage() {
	instIds := ws.getInstIds()
	for _, v := range ws.SubscribeChannel {
		ws.SendStructMessage(OpSubscribe, v, instIds)
	}
}

// SendStructMessage 发送信息
func (ws *WsInstance) SendStructMessage(op string, channel Channel, instIds []string) {
	message := &OkxParams{Args: make([]*Arg, 0), Op: op}
	for _, v := range instIds {
		message.Args = append(message.Args, &Arg{
			Channel: channel,
			InstID:  v,
		})
	}

	ws.SendMessage(utils.JsonToBytes(message))
}

// getInstIds 获取需要订阅的消息。
func (ws *WsInstance) getInstIds() []string {
	var instIds []string
	if result := database.Db.Model(&models.Product{}).
		Where("admin_id = ?", models.SuperAdminId).
		Where("type = ?", models.ProductTypeOkex).
		Where("status = ?", models.ProductStatusActivate).
		Pluck("symbol", &instIds); result.Error != nil {
		logger.Logger.Error(logger.LogMsgOkx, zap.String("method", "getInstIds"), zap.Error(result.Error))
		return nil
	}
	return instIds
}

// GetMaxReconnect 获取重连次数
func (ws *WsInstance) GetMaxReconnect() int {
	return ws.MaxReconnect
}

// SetMaxReconnect 设置重连次数
func (ws *WsInstance) SetMaxReconnect(MaxReconnect int) *WsInstance {
	ws.MaxReconnect = MaxReconnect
	return ws
}

// GetServerAddr 获取服务地址
func (ws *WsInstance) GetServerAddr() string {
	return ws.ServerAddr
}

// SetServerAddr 设置服务地址
func (ws *WsInstance) SetServerAddr(ServerAddr string) *WsInstance {
	ws.ServerAddr = ServerAddr
	return ws
}

// GetSubscribeChannel 获取订阅通道
func (ws *WsInstance) GetSubscribeChannel() []Channel {
	return ws.SubscribeChannel
}

// SetSubscribeChannel 添加订阅的通道
func (ws *WsInstance) SetSubscribeChannel(SubscribeChannel ...Channel) *WsInstance {
	temp := make(map[Channel]Channel)
	for _, v := range ws.SubscribeChannel {
		temp[v] = v
	}
	for _, channel := range SubscribeChannel {
		if _, ok := temp[channel]; !ok {
			ws.SubscribeChannel = append(ws.SubscribeChannel, channel)
		}
	}
	return ws
}

// GetSubscribeKlineChannel 获取 Kline 通道
func (ws *WsInstance) GetSubscribeKlineChannel() []Channel {
	channelKlineList := make([]Channel, 0)
	for k, _ := range ChannelKlineMap {
		channelKlineList = append(channelKlineList, k)
	}
	return channelKlineList
}

// SetSubscribeKlineChannel 设置 Kline 通道
func (ws *WsInstance) SetSubscribeKlineChannel() *WsInstance {
	channelKlineList := make([]Channel, 0)
	for k, _ := range ChannelKlineMap {
		channelKlineList = append(channelKlineList, k)
	}
	ws.SubscribeChannel = channelKlineList
	return ws
}

// SetSubscribeDefaultChannel 设置 默认 通道
func (ws *WsInstance) SetSubscribeDefaultChannel() *WsInstance {
	channelDefaultList := make([]Channel, 0)
	for k, _ := range ChannelMap {
		channelDefaultList = append(channelDefaultList, k)
	}
	ws.SubscribeChannel = channelDefaultList

	return ws
}

// ClearSubscribeChannel 获取订阅通道
func (ws *WsInstance) ClearSubscribeChannel() *WsInstance {
	ws.SubscribeChannel = []Channel{}
	return ws
}
