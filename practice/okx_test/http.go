package okx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
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

// RspData okx返回数据
type RspData struct {
	Code string      //	返回代码
	Msg  string      //	消息
	Data interface{} //	内容
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
		err = errors.New(fmt.Sprintln("code = ", result.StatusCode))
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return nil, err
	}

	//	接口是否正常
	respData := new(RspData)
	err = json.Unmarshal(body, &respData)
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	}

	if respData.Code != "0" {
		err = errors.New(respData.Msg)
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return nil, err
	}

	// 缓存1s数据，防止过量请求
	respDataBytes, _ := json.Marshal(&respData.Data)
	if _, err = rds.Do("SETEX", currentRestURL, 5, respDataBytes); err != nil {
		return nil, err
	}

	return respDataBytes, nil
}
