package okx

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gotest/common/module/cache"
	"gotest/common/module/logs"
	"gotest/common/utils"
	"net/url"
)

const ()

// BooksData 深度内部数据
type BooksData struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Ts        string     `json:"ts"`
	Checksum  int        `json:"checksum"`
	PrevSeqId int        `json:"prevSeqId"`
	SeqId     int        `json:"seqId"`
}

// books 返回的推送数据
type books struct {
	Arg    Arg         `json:"Arg"`
	Action string      `json:"action"`
	Data   []BooksData `json:"data"`
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

// GetWsBooks 获取深度行情数据
func GetWsBooks(instId string) (*BooksData, error) {
	rdsConn := cache.RdsPool.Get()
	defer rdsConn.Close()
	result, err := redis.Bytes(rdsConn.Do("HGET", ChannelMap[ChannelBooks], instId))
	if err != nil {
		logs.Logger.Error("okx", zap.Error(err))
		return nil, err
	}

	// 解析redis 数据
	data := &BooksData{}
	_ = json.Unmarshal(result, &data)

	return data, nil
}

// GetWsBooksString 获取深度行情字符串数据
func GetWsBooksString(instId string) (string, error) {
	data, err := GetWsTicker(instId)
	if err != nil {
		return "", err
	}

	dataString := utils.ObjToString(data)
	return dataString, nil
}
