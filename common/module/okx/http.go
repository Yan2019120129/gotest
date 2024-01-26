package okx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/frame/my_frame/module/cache"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	serverOkxAddr     = "https://www.okx.com/api/v5/market/"
	serverAddrTrades  = "trades"
	serverAddrBooks   = "books"
	serverAddrCandles = "candles"
)

var serverOkxAddrMap = map[string]string{
	serverAddrTrades:  serverOkxAddr + serverAddrTrades,
	serverAddrBooks:   serverOkxAddr + serverAddrBooks,
	serverAddrCandles: serverOkxAddr + serverAddrCandles,
}

// RspData okx返回数据
type RspData struct {
	Code string      //	返回代码
	Msg  string      //	消息
	Data interface{} //	内容
}

// KlineData k线图数据
type KlineData struct {
	OpenPrice  float64 `json:"openPrice"`  //开盘价格
	HighPrice  float64 `json:"highPrice"`  //最高价格
	LowsPrice  float64 `json:"lowsPrice"`  //最低价格
	ClosePrice float64 `json:"closePrice"` //收盘价格
	Vol        float64 `json:"vol"`        //交易量
	Amount     float64 `json:"amount"`     //成交额
	CreatedAt  int64   `json:"createdAt"`  //开盘时间
}

// TradesData 交易数据
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

// GetTrades 交易量
func GetTrades(instId string) ([]*TradesData, error) {
	// 发送请求获取交易深度数据
	query := url.Values{"instId": {instId}}
	resp, err := Get(serverOkxAddrMap[serverAddrTrades], query)
	if err != nil {
		return nil, err
	}

	data := make([]*TradesData, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}

// GetBooks 获取交易深度数据
func GetBooks(instId string) ([]*BooksData, error) {
	// 发送请求获取交易深度数据
	query := url.Values{"instId": {instId}, "sz": {"60"}}
	resp, err := Get(serverOkxAddrMap[serverAddrBooks], query)
	if err != nil {
		return nil, err
	}

	data := make([]*BooksData, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}

// GetKline 获取k线图数据
func GetKline(instId, bar string) ([]*KlineData, error) {
	query := url.Values{"instId": {instId}, "limit": {"300"}, "bar": {bar}}
	resp, err := Get(serverOkxAddrMap[serverAddrCandles], query)
	if err != nil {
		return nil, err
	}

	data := make([]*KlineData, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}

// Get 发送请求公共方法
func Get(path string, query url.Values) ([]byte, error) {
	rds := cache.RdsPool.Get()
	defer rds.Close()
	currentRestURL := path + "?" + query.Encode()

	// 判断是否存在数据
	if data, err := redis.Bytes(rds.Do("get", currentRestURL)); err == nil {
		return data, nil
	}

	// 发送请求获取数据
	result, err := http.Get(currentRestURL)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	// 读取body数据
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	// 判断是否状态200
	if result.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintln("code = ", result.StatusCode))
	}

	//	接口是否正常
	respData := new(RspData)
	_ = json.Unmarshal(body, &respData)
	if respData.Code != "0" {
		return nil, errors.New(respData.Msg)
	}

	// 缓存1s数据，防止过量请求
	respDataBytes, _ := json.Marshal(&respData.Data)
	if _, err = rds.Do("SETEX", currentRestURL, 5, respDataBytes); err != nil {
		return nil, err
	}

	return respDataBytes, nil
}

// GetAll 获取全部的Okx 数据
func GetAll() {
	for {
		books, err := GetBooks("BTC-USDT")
		if err != nil {
			log.Println("booksErr", err)
		}
		log.Println("books", books)

		kline, err := GetKline("BTC-USDT", "3m")
		if err != nil {
			log.Println("booksErr", err)
		}
		log.Println("kline", kline)

		Trades, err := GetTrades("BTC-USDT")
		if err != nil {
			log.Println("booksErr", err)
		}
		log.Println("Trades", Trades)
		time.Sleep(3 * time.Second)
	}
}
