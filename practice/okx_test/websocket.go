package okx_test

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"gotest/common/module/cache"
	"gotest/common/module/gorm/database"
	"gotest/common/utils"
	"gotest/frame/my-fiber/models"
	"log"
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

// tickerData 返回的推送数据
type tickerData struct {
	Arg  arg          `json:"arg"`
	Data []TickerData `json:"data"`
}

// tradesData 返回的推送数据
type tradesData struct {
	Arg  arg          `json:"arg"`
	Data []TradesData `json:"data"`
}

// booksData 返回的推送数据
type booksData struct {
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

// TradesData 交易推送内部数据
type TradesData struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Ts      string `json:"ts"`
	Count   string `json:"count"`
}

// BooksData 深度内部数据
type BooksData struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Ts        string     `json:"ts"`
	Checksum  int        `json:"checksum"`
	PrevSeqId int        `json:"prevSeqId"`
	SeqId     int        `json:"seqId"`
}

// ProductData 产品数据
type ProductData struct {
	Id   int    // 产品id
	Name string // 产品名
}

// ConnectWS 连接okx websocket。
func (ws *WsInstance) connect() (err error) {
	// 添加协程使用WaitGroup管理线程状态
	ws.Conn, _, err = websocket.DefaultDialer.Dial(ws.ServerAddr, nil)
	if err != nil {
		log.Fatal("WebSocket连接错误:", err)
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
	ws.send()

	// 心跳检测
	go ws.heartbeat()

	// 读取信息
	go ws.readMessages()

	// 处理并发送信息
	go ws.handlePublish()

	// 订阅信息
	go ws.subscribeMessages()

	ws.GetTicker("BTC-USDT")

	//// 定时15秒后关闭
	//go ws.Close()
}

// Close 关闭链接
func (ws *WsInstance) Close() {
	if err := ws.Conn.Close(); err != nil {
		log.Println("okx warn:", err)
		return
	}
	log.Println("okx cloned。")
}

// sendMessages 发送消息。
func (ws *WsInstance) sendMessages(message []byte) {
	if err := ws.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Println("Write error", err)
	}
}

// ReadMessages 读取消息。
func (ws *WsInstance) readMessages() {
	defer fmt.Println("关闭：readMessages")
	for {
		// 当读取消息发送错误，则重新连接
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			log.Println("reading error", err)
			ws.reconnect()
			continue
		}
		ws.data <- string(message)
	}
}

// handlePublish 处理发布信息。
func (ws *WsInstance) handlePublish() {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()
	defer fmt.Println("关闭：handlePublish")
	for {
		message := <-ws.data
		// 映射到指定的频道
		tickerMessage, tradesMessage, bookMessage := new(tickerData), new(tradesData), new(booksData)
		switch {
		// 将行情数据发布到redis tickers 通道
		case strings.Contains(message, ChannelTicker) && strings.Contains(message, "data"):
			if err := json.Unmarshal([]byte(message), &tickerMessage); err == nil {
				_, _ = rdsConn.Do("PUBLISH", tickerMessage.Arg.Channel, tickerMessage.Data)
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

				reply, err := rdsConn.Do("HSET", ChannelTicker, tickerMessage.Data[0].InstId, utils.ObjToByteList(data))
				if err != nil {
					log.Println("错误", err)
				}
				log.Println("setRdsData", data)
				log.Println("setRds", reply)
			}

		// 将行情数据发布到redis books 通道
		case (strings.Contains(message, ChannelBooks) || strings.Contains(message, ChannelBooks5) || strings.Contains(message, ChannelBooksL2Tbt) || strings.Contains(message, ChannelBooks50L2Tbt)) && strings.Contains(message, "data"):
			if err := json.Unmarshal([]byte(message), &bookMessage); err == nil {
				_, _ = rdsConn.Do("PUBLISH", bookMessage.Arg.Channel, bookMessage.Data)
			}

		// 将行情数据发布到redis trades 通道
		case strings.Contains(message, ChannelTrades) && strings.Contains(message, "data"):
			if err := json.Unmarshal([]byte(message), &tradesMessage); err == nil {
				_, _ = rdsConn.Do("PUBLISH", tradesMessage.Arg.Channel, tradesMessage.Data)
			}

		// 打印错误信息
		case strings.Contains(message, "error") && strings.Contains(message, "msg"):
			log.Println("error", message)

		// 打印其他返回信息
		default:
			log.Println("other", message)
		}
	}
}

// HandleMessages 每个几秒尝试发送连接消息，检测连接心跳
func (ws *WsInstance) heartbeat() {
	defer fmt.Println("关闭：heartbeat")
	for {

		// 发送心跳消息
		ws.sendMessages([]byte("ping"))

		// 每隔三秒发送消息
		time.Sleep(5 * time.Second)
	}
}

// reconnect 重新连接
func (ws *WsInstance) reconnect() {
	// 根觉最大的连接次数，断开连接的时候重新连接
	for i := 0; i < ws.MaxReconnect; i++ {
		log.Println("重新连接:", strconv.Itoa(ws.MaxReconnect))

		if err := ws.connect(); err != nil {
			continue
		}
		ws.send()
		return
	}
}

// subscribeMessages 订阅消息。
func (ws *WsInstance) subscribeMessages() {
	rdsConn := cache.RdsPubSubConn
	defer rdsConn.Close()
	defer log.Println("关闭：subscribeMessages")
	for {
		if err := rdsConn.PSubscribe(ChannelTicker, func(data []byte) {
			log.Println(ChannelTicker, string(data))
		}); err != nil {
			log.Println("cache error:"+ChannelTicker, err)
		}

		if err := rdsConn.Subscribe(ChannelBooks, func(data []byte) {
			log.Println(ChannelBooks, string(data))
		}); err != nil {
			log.Println("cache error:"+ChannelBooks, err)
		}

		if err := rdsConn.Subscribe(ChannelTrades, func(data []byte) {
			log.Println(ChannelTrades, string(data))
		}); err != nil {
			log.Println("cache error:"+ChannelTicker, err)
		}

		// 打印错误
		if err := rdsConn.Subscribe("error", func(data []byte) {
			log.Println("error", string(data))
		}); err != nil {
			log.Println("cache error: error", err)
		}

		// 打印其他信息，
		if err := rdsConn.Subscribe("other", func(data []byte) {
			log.Println("other:", string(data))
		}); err != nil {
			log.Println("cache error: other", err)
		}
	}
}

// send 发送测试数据
func (ws *WsInstance) send() {
	tickerMessage := &TickerParams{Args: make([]*arg, 0), Op: OpSubscribe}
	tradesMessage := &TickerParams{Args: make([]*arg, 0), Op: OpSubscribe}
	bookMessage := &TickerParams{Args: make([]*arg, 0), Op: OpSubscribe}

	productList := ws.getSendMessage()
	for _, v := range productList {
		v.Name = strings.ReplaceAll(v.Name, "/", "-")
		tickerMessage.Args = append(tickerMessage.Args, &arg{
			Channel: ChannelTicker,
			InstID:  v.Name,
		})
		tradesMessage.Args = append(tradesMessage.Args, &arg{
			Channel: ChannelTrades,
			InstID:  v.Name,
		})
		bookMessage.Args = append(bookMessage.Args, &arg{
			Channel: ChannelBooks,
			InstID:  v.Name,
		})
	}
	ws.sendMessages(utils.ObjToByteList(tickerMessage))
	ws.sendMessages(utils.ObjToByteList(tradesMessage))
	ws.sendMessages(utils.ObjToByteList(bookMessage))
}

// getSendMessage 获取需要订阅的消息。
func (ws *WsInstance) getSendMessage() []*ProductData {
	productList := make([]*ProductData, 0)
	if result := database.DB.Model(&models.Product{}).
		Select("id", "name").
		Where("type = ?", models.ProductTypeOkex).
		Where("status = ?", models.ProductStatusActivate).
		Find(&productList); result.Error != nil {
		return nil
	}
	return productList
}

// GetTicker 获取行情数据
func (ws *WsInstance) GetTicker(instId string) (*TickerRdsData, error) {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()
	result, err := redis.Bytes(rdsConn.Do("HGET", ChannelTicker, instId))
	if err != nil {
		log.Println("rdsResultErr:", err)
	}
	tickerRdsData := &TickerRdsData{}
	_ = json.Unmarshal(result, &tickerRdsData)
	log.Println("rdsResult:", tickerRdsData)
	return tickerRdsData, nil
}
