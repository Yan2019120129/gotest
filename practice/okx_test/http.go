package okx_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/common/module/cache"
	"gotest/common/module/log/zap_log"
	"io"
	"net/http"
	"net/url"
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

type RspData struct {
	Code string      //	返回代码
	Msg  string      //	消息
	Data interface{} //	内容
}

type KlineData struct {
	OpenPrice  float64 `json:"openPrice"`  //开盘价格
	HighPrice  float64 `json:"highPrice"`  //最高价格
	LowsPrice  float64 `json:"lowsPrice"`  //最低价格
	ClosePrice float64 `json:"closePrice"` //收盘价格
	Vol        float64 `json:"vol"`        //交易量
	Amount     float64 `json:"amount"`     //成交额
	CreatedAt  int64   `json:"createdAt"`  //开盘时间
}

// GetTrades 交易量
func GetTrades(instId string) ([]*TradesData, error) {
	// 发送请求获取交易深度数据
	query := url.Values{"instId": {instId}, "limit": {"100"}}
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
	query := url.Values{"instId": {instId}, "limit": {"100"}, "bar": {bar}}
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
	data, err := redis.Bytes(rds.Do("get", currentRestURL))
	if err == nil {
		//logger.Logger.Info("info", zap.Reflect("data", data))
		return data, nil
	}

	// 发送请求获取数据
	result, err := http.Get(currentRestURL)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	zap_log.Logger.Info("A request has been sent")

	// 读取body数据
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	// 判断是否状态200
	if result.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintln("code = ", result.StatusCode))
	}

	return handle(currentRestURL, body)
}

// handle 处理返回的数据
func handle(kay string, data []byte) ([]byte, error) {
	rds := cache.RdsPool.Get()
	defer rds.Close()

	//	接口是否正常
	respData := new(RspData)
	_ = json.Unmarshal(data, &respData)
	if respData.Code != "0" {
		return nil, errors.New(respData.Msg)
	}

	respDataBytes, _ := json.Marshal(respData.Data)

	// 缓存1s数据，防止过量请求
	if _, err := rds.Do("SETEX", kay, 5, respDataBytes); err != nil {
		return nil, err
	}

	return respDataBytes, nil
}
