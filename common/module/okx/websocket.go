package okx

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gotest/frame/my-fiber/models"
	"gotest/frame/my_frame/module/cache"
	"gotest/frame/my_frame/module/gorm/database"
	"gotest/frame/my_frame/module/logs"
	"gotest/frame/my_frame/utils"
	"strconv"
	"strings"
	"time"
)

// ServerOkxAddr 产品行情地址。
const (
	// ServerOkxAddr okx 行情websocket 地址
	//ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"
	ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"

	// ServerCandleAndTradeAddr okx 行业websocket 地址
	ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"

	// OpSubscribe 订阅
	OpSubscribe = "subscribe"

	// OpUnsubscribe 取消订阅
	OpUnsubscribe = "unsubscribe"

	// ChannelTicker 行情频道
	ChannelTicker = "tickers"

	// ChannelBooks 深度频道
	ChannelBooks = "books"

	// ChannelBooks5 深度频道
	ChannelBooks5 = "books5"

	// ChannelBooksL2Tbt 深度频道
	ChannelBooksL2Tbt = "books-l2-tbt"

	// ChannelBooks50L2Tbt 深度频道
	ChannelBooks50L2Tbt = "books50-l2-tbt"

	// ChannelTrades 交易频道
	ChannelTrades = "trades"
)

// Instance 行情 websocket 实例
var Instance = &WsInstance{
	Conn:         new(websocket.Conn),    // websocket 实例
	MaxReconnect: 5,                      // 最大重连次数
	ServerAddr:   ServerOkxAddr,          // 订阅的通道，用于使用指定的连接地址
	data:         make(chan string, 100), // 数据传输通道，用户接收处理数据
}

// WsInstance websocket实例。
type WsInstance struct {
	Conn         *websocket.Conn
	MaxReconnect int         // 设置最大重连次数
	ServerAddr   string      // 连接地址
	data         chan string // 用于接收，处理数据。
}

// TickerParams 发送参数。
type TickerParams struct {
	Op   string `json:"op"`   // 操作，subscribe unsubscribe
	Args []*arg `json:"args"` // 请求订阅的频道列表
}

// ticker 返回的推送数据
type ticker struct {
	Arg  arg          `json:"arg"`
	Data []TickerData `json:"data"`
}

// trades 返回的推送数据
type trades struct {
	Arg  arg          `json:"arg"`
	Data []TradesData `json:"data"`
}

// books 返回的推送数据
type books struct {
	Arg    arg         `json:"arg"`
	Action string      `json:"action"`
	Data   []BooksData `json:"data"`
}

// Arg 币种订阅频道。
type arg struct {
	Channel string `json:"channel"` // 订阅的通道
	InstID  string `json:"instId"`  // 货币类型
}

// TickerData 行情推送的数据
type TickerData struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        string `json:"ts"`
}

// TickerRdsData 用于存储到rds
type TickerRdsData struct {
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	Ts        string `json:"ts"`
}

// TickerConvertData 用于存储到用户外部方便使用
type TickerConvertData struct {
	InstId    string  `json:"instId"`
	Last      float64 `json:"last"`
	LastSz    float64 `json:"lastSz"`
	Open24h   float64 `json:"open24h"`
	High24h   float64 `json:"high24h"`
	Low24h    float64 `json:"low24h"`
	VolCcy24h float64 `json:"volCcy24h"`
	Vol24h    float64 `json:"vol24h"`
	Ts        int     `json:"ts"`
}

// ProductData 产品数据
type ProductData struct {
	Id     int    // 产品id
	Symbol string // 标识
}

// ConnectWS 连接okx websocket。
func (ws *WsInstance) connect() (err error) {
	// 添加协程使用WaitGroup管理线程状态
	ws.Conn, _, err = websocket.DefaultDialer.Dial(ws.ServerAddr, nil)
	if err != nil {
		logs.Logger.Error("okx", zap.String("method", "connect"), zap.Error(err))
		return err
	}

	return nil
}

// Run 启动websocket
func (ws *WsInstance) Run() {
	// 链接websocket
	if err := ws.connect(); err != nil {
		return
	}

	// 发送测试数据
	ws.sendAll()

	// 心跳检测
	go ws.heartbeat()

	// 读取信息
	go ws.read()

	// 处理并发送信息
	go ws.handlePublish()

	// 订阅信息
	go ws.subscribe()

	//_, _ = GetRdsTicker("BTC-USDT")
}

// Close 关闭链接
func (ws *WsInstance) Close() {
	if err := ws.Conn.Close(); err != nil {
		logs.Logger.Error("okx", zap.String("method", "Close"), zap.Error(err))
		return
	}
}

// send 发送消息。
func (ws *WsInstance) send(message []byte) {
	if err := ws.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		logs.Logger.Error("okx", zap.String("method", "send"), zap.Error(err))
	}
}

// ReadMessages 读取消息。
func (ws *WsInstance) read() {
	defer fmt.Println("关闭：read")
	for {
		// 当读取消息发送错误，则重新连接
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.reconnect()
			logs.Logger.Error("okx", zap.String("method", "read"), zap.Error(err))
			continue
		}
		ws.data <- string(message)
	}
}

// handlePublish 处理发布信息。
func (ws *WsInstance) handlePublish() {
	rds := cache.RdsPool.Get()
	defer rds.Close()
	defer logs.Logger.Info("okx", zap.String("method", "handlePublish"), zap.String("info", "handle publish close"))
	for {
		message := <-ws.data
		// 映射到指定的频道
		tickerMessage, tradesMessage, bookMessage := new(ticker), new(trades), new(books)
		switch {
		// 将行情数据发布到redis tickers 通道
		case strings.Contains(message, ChannelTicker) && strings.Contains(message, "data"):
			if err := json.Unmarshal([]byte(message), &tickerMessage); err == nil {
				if _, err = rds.Do("PUBLISH", tickerMessage.Arg.Channel, tickerMessage.Data); err != nil {
					logs.Logger.Error("okx", zap.String("method", "handlePublish"), zap.Error(err))
				}
				data := &TickerRdsData{
					InstId:    tickerMessage.Data[0].InstId,
					Last:      tickerMessage.Data[0].Last,
					LastSz:    tickerMessage.Data[0].LastSz,
					Open24h:   tickerMessage.Data[0].Open24h,
					High24h:   tickerMessage.Data[0].High24h,
					Low24h:    tickerMessage.Data[0].Low24h,
					VolCcy24h: tickerMessage.Data[0].VolCcy24h,
					Vol24h:    tickerMessage.Data[0].Vol24h,
					Ts:        tickerMessage.Data[0].Ts,
				}

				if _, err = rds.Do("HSET", ChannelTicker, tickerMessage.Data[0].InstId, utils.ObjToByteList(data)); err != nil {
					logs.Logger.Error("okx", zap.String("method", "handlePublish"), zap.Error(err))
				}
			}

		// 将行情数据发布到redis books 通道
		case (strings.Contains(message, ChannelBooks) || strings.Contains(message, ChannelBooks5) || strings.Contains(message, ChannelBooksL2Tbt) || strings.Contains(message, ChannelBooks50L2Tbt)) && strings.Contains(message, "data"):
			if err := json.Unmarshal([]byte(message), &bookMessage); err == nil {
				if _, err = rds.Do("PUBLISH", bookMessage.Arg.Channel, bookMessage.Data); err != nil {
					logs.Logger.Error("okx", zap.String("method", "handlePublish"), zap.Error(err))
				}
			}

		// 将行情数据发布到redis trades 通道
		case strings.Contains(message, ChannelTrades) && strings.Contains(message, "data"):
			if err := json.Unmarshal([]byte(message), &tradesMessage); err == nil {
				if _, err = rds.Do("PUBLISH", tradesMessage.Arg.Channel, tradesMessage.Data); err != nil {
					logs.Logger.Error("okx", zap.String("method", "handlePublish"), zap.Error(err))
				}
			}

		// 打印错误信息
		case strings.Contains(message, "error") && strings.Contains(message, "msg"):
			logs.Logger.Error("okx", zap.String("method", "handlePublish"), zap.String("error", message))

		// 打印其他返回信息
		default:
			logs.Logger.Error("okx", zap.String("method", "handlePublish"), zap.String("other", message))
		}
	}
}

// heartbeat 每个几秒尝试发送连接消息，检测连接心跳
func (ws *WsInstance) heartbeat() {
	defer logs.Logger.Info("okx", zap.String("method", "heartbeat"), zap.String("info", "heartbeat close"))
	for {

		// 发送心跳消息
		ws.send([]byte("ping"))

		// 每隔三秒发送消息
		time.Sleep(5 * time.Second)
	}
}

// reconnect 重新连接
func (ws *WsInstance) reconnect() {
	defer logs.Logger.Info("okx", zap.String("method", "reconnect"), zap.String("info", "reconnect close"))
	// 根觉最大的连接次数，断开连接的时候重新连接
	for i := 0; i < ws.MaxReconnect; i++ {
		logs.Logger.Warn("okx", zap.String("method", "reconnect"), zap.Int("sum", i))

		if err := ws.connect(); err != nil {
			logs.Logger.Error("okx", zap.String("method", "reconnect"), zap.Int("sum", i), zap.Error(err))
			continue
		}
		ws.sendAll()
		return
	}
}

// subscribe 订阅消息。
func (ws *WsInstance) subscribe() {
	defer logs.Logger.Info("okx", zap.String("method", "subscribe"), zap.String("info", "subscribe close"))
	rds := cache.RdsPubSubConn
	defer rds.Close()
	for {
		if err := rds.Subscribe(ChannelTicker, func(data []byte) {
			logs.Logger.Info("okx", zap.String("method", "subscribe"), zap.String("channel", ChannelTicker), zap.ByteString("data", data))
		}); err != nil {
			logs.Logger.Error("okx", zap.String("method", "subscribe"), zap.String("channel", ChannelTicker), zap.Error(err))
		}

		if err := rds.Subscribe(ChannelBooks, func(data []byte) {
			logs.Logger.Info("okx", zap.String("method", "subscribe"), zap.String("channel", ChannelBooks), zap.ByteString("data", data))
		}); err != nil {
			logs.Logger.Error("okx", zap.String("method", "subscribe"), zap.String("channel", ChannelBooks), zap.Error(err))
		}

		if err := rds.Subscribe(ChannelTrades, func(data []byte) {
			logs.Logger.Info("okx", zap.String("method", "subscribe"), zap.String("channel", ChannelTrades), zap.ByteString("data", data))
		}); err != nil {
			logs.Logger.Error("okx", zap.String("method", "subscribe"), zap.String("channel", ChannelTrades), zap.Error(err))
		}
	}
}

// sendAll 发送测试数据
func (ws *WsInstance) sendAll() {
	tickerMessage := &TickerParams{Args: make([]*arg, 0), Op: OpSubscribe}
	tradesMessage := &TickerParams{Args: make([]*arg, 0), Op: OpSubscribe}
	bookMessage := &TickerParams{Args: make([]*arg, 0), Op: OpSubscribe}

	productList := ws.getInstIds()
	for _, v := range productList {
		tickerMessage.Args = append(tickerMessage.Args, &arg{
			Channel: ChannelTicker,
			InstID:  v.Symbol,
		})
		//tradesMessage.Args = append(tradesMessage.Args, &arg{
		//	Channel: ChannelTrades,
		//	InstID:  v.Symbol,
		//})
		//bookMessage.Args = append(bookMessage.Args, &arg{
		//	Channel: ChannelBooks,
		//	InstID:  v.Symbol,
		//})
	}
	ws.send(utils.ObjToByteList(tickerMessage))
	ws.send(utils.ObjToByteList(tradesMessage))
	ws.send(utils.ObjToByteList(bookMessage))
}

// getInstIds 获取需要订阅的消息。
func (ws *WsInstance) getInstIds() []*ProductData {
	productList := make([]*ProductData, 0)
	if result := database.DB.Model(&models.Product{}).
		Select("id", "symbol").
		Where(1).
		Where("type = ?", models.ProductTypeOkex).
		Where("status = ?", models.ProductStatusActivate).
		Find(&productList); result.Error != nil {
		logs.Logger.Error("okx", zap.String("method", "getInstIds"), zap.Error(result.Error))
		return nil
	}
	return productList
}

// GetRdsTicker 获取行情数据未转换类型
func GetRdsTicker(instId string) (*TickerRdsData, error) {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()
	result, err := redis.Bytes(rdsConn.Do("HGET", ChannelTicker, instId))
	if err != nil {
		logs.Logger.Error("okx", zap.String("method", "GetRdsTicker"), zap.Error(err))
		return nil, err
	}

	tickerRdsData := &TickerRdsData{}
	_ = json.Unmarshal(result, &tickerRdsData)

	return tickerRdsData, nil
}

// GetTicker 获取行情数据
func GetTicker(instId string) (*TickerConvertData, error) {
	data, err := GetRdsTicker(instId)
	if err != nil {
		return nil, err
	}
	tickerConvertData := &TickerConvertData{}
	tickerConvertData.LastSz, _ = strconv.ParseFloat(data.LastSz, 64)
	tickerConvertData.Open24h, _ = strconv.ParseFloat(data.Open24h, 64)
	tickerConvertData.High24h, _ = strconv.ParseFloat(data.High24h, 64)
	tickerConvertData.Low24h, _ = strconv.ParseFloat(data.Low24h, 64)
	tickerConvertData.VolCcy24h, _ = strconv.ParseFloat(data.VolCcy24h, 64)
	tickerConvertData.Vol24h, _ = strconv.ParseFloat(data.Vol24h, 64)
	ts, _ := strconv.ParseInt(data.Ts, 63, 10)
	tickerConvertData.Ts = int(ts)
	return tickerConvertData, nil
}
