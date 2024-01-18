package okex

import (
	"encoding/json"
	"errors"
	"fmt"
	"gotest/common/module/cache"
	"gotest/common/module/log/zap_log"
	"gotest/practice/okex/dto"
	"gotest/practice/okex/utils"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// SocketURL websocket 连接URL
const SocketURL = "wss://ws.okx.com:8443/ws/v5/public"
const RestURL = "https://www.okx.com"

var Okex *OkexStruct

type OkexStruct struct {
	conn           *websocket.Conn                  //	websocket客户端
	sync           sync.RWMutex                     //	map 加锁操作
	CurrentTickers map[string]*dto.ProductDataAttrs //	当前行情最新价格
}

type OkexResp struct {
	Code string      `json:"code"` //	返回代码
	Msg  string      `json:"msg"`  //	消息
	Data interface{} `json:"data"` //	内容
}

func NewOkexStruct() *OkexStruct {
	return &OkexStruct{
		conn:           nil,
		CurrentTickers: map[string]*dto.ProductDataAttrs{},
	}
}

// Get GET请求
func (_OkexStruct *OkexStruct) Get(path string, params map[string]interface{}) ([]byte, error) {
	paramsStr := utils.MapBuildQuery(params)
	currentRestURL := RestURL + path
	if paramsStr != "" {
		currentRestURL = currentRestURL + "?" + paramsStr
	}

	//	1s缓存数据, 防止过量请求
	rds := cache.RdsPool.Get()
	defer rds.Close()

	// 判断是否存在数据
	replyByte, err := redis.Bytes(rds.Do("GET", currentRestURL))
	if err == nil {
		return replyByte, nil
	}

	resp, err := http.Get(currentRestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 判断是否状态200
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintln("code = ", resp.StatusCode))
	}

	// 读取body数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//	接口是否正常
	respData := new(OkexResp)
	_ = json.Unmarshal(body, &respData)
	if respData.Code != "0" {
		return nil, errors.New(respData.Msg)
	}

	respDataBytes, _ := json.Marshal(respData.Data)

	// 缓存1s数据，防止过量请求
	_, _ = rds.Do("SETEX", currentRestURL, 1, respDataBytes)
	return respDataBytes, nil
}

// ForReader 循环读取消息
func (_OkexStruct *OkexStruct) Reader() {
	if _OkexStruct.conn == nil {
		//	重新连接
		_OkexStruct.ReconnectOkex()
		return
	}

	defer _OkexStruct.Colse()

	for {
		_, messageData, err := _OkexStruct.conn.ReadMessage()
		if err != nil {
			zap_log.Logger.Debug("okex ====> 连接已关闭~  5s后进行重新连接....", zap.String("data", err.Error()))

			//	重新连接
			_OkexStruct.ReconnectOkex()
			break
		}

		//	传播数据给客户端
		data := new(SubscribeData)
		err = json.Unmarshal(messageData, &data)
		if err == nil && data.Arg != nil {
			//	保存当前行情数据
			if data.Arg.Channel == "tickers" {
				currentTickers := make([]*Ticker, 0)
				tickersBytes, _ := json.Marshal(data.Data)
				err = json.Unmarshal(tickersBytes, &currentTickers)
				if err == nil && len(currentTickers) > 0 {
					_OkexStruct.SetTicker(data.Arg.InstId, &dto.ProductDataAttrs{
						InstId:    currentTickers[0].InstId,
						Last:      currentTickers[0].GetLast(),
						LastSz:    currentTickers[0].GetLastSz(),
						Open24h:   currentTickers[0].GetOpen24h(),
						High24h:   currentTickers[0].GetHigh24h(),
						Low24h:    currentTickers[0].GetLow24h(),
						Vol24h:    currentTickers[0].GetVol24h(),
						Amount24h: currentTickers[0].GetAmount24h(),
						Ts:        currentTickers[0].GetTs(),
					})
				}
			}
			zap_log.Logger.Info("info", zap.Reflect("okx:", data))
			Socket.WriterAllClient(data)
		}
	}
}

// connect 连接
func (_OkexStruct *OkexStruct) Connect() {
	conn, _, err := websocket.DefaultDialer.Dial(SocketURL, nil)
	if err != nil {
		fmt.Println("连接 okx websocket 失败!", err.Error())
		return
	}
	_OkexStruct.conn = conn
}

// Colse 关闭连接
func (_OkexStruct *OkexStruct) Colse() {
	_ = _OkexStruct.conn.Close()
	_OkexStruct.conn = nil
}

// Writer 订阅数据
func (_OkexStruct *OkexStruct) Subscribe(msg *Subscribe) {
	if _OkexStruct.conn != nil {
		zap_log.Logger.Info("info", zap.Reflect("message", msg))
		_ = _OkexStruct.conn.WriteJSON(msg)
	}
}

// GetTicker 获取币种行情
func (_OkexStruct *OkexStruct) GetTicker(symbol string) *dto.ProductDataAttrs {
	_OkexStruct.sync.RLock()
	defer _OkexStruct.sync.RUnlock()

	if _, ok := _OkexStruct.CurrentTickers[symbol]; ok {
		return _OkexStruct.CurrentTickers[symbol]
	}
	return nil
}

// SetTicker 设置币种行情
func (_OkexStruct *OkexStruct) SetTicker(symbol string, attrs *dto.ProductDataAttrs) {
	_OkexStruct.sync.Lock()
	defer _OkexStruct.sync.Unlock()
	_OkexStruct.CurrentTickers[symbol] = attrs
}

// TickerUpdatePrice 定时任务更新最新行情数据
func (_OkexStruct *OkexStruct) TickerUpdatePrice() {
	_OkexStruct.sync.RLock()
	for symbol, ticker := range _OkexStruct.CurrentTickers {
		tickerBytes, _ := json.Marshal(ticker)
		zap_log.Logger.Info("info", zap.Reflect("update", symbol))
		zap_log.Logger.Info("info", zap.Reflect("TickerData", tickerBytes))
		//symbol = strings.ReplaceAll(symbol, "-", "/")
		//if result := database.DB.Where("type = ?", dto.ProductTypeOkex).Where("name = ?").Update("data", string(tickerBytes)); result.Error != nil {
		//	logger.Logger.Warn("warn", zap.Error(result.Error))
		//}
	}
	_OkexStruct.sync.RUnlock()
}

// ReconnectOkex 重新连接okex
func (_OkexStruct *OkexStruct) ReconnectOkex() {
	time.AfterFunc(5*time.Second, func() {
		_OkexStruct.Connect()
		_OkexStruct.SubscribeTickers().Reader()
	})
}
