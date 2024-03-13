package http_t

import (
	"errors"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"gotest/common/utils"
	"gotest/practice/http_t/dto"
	"io"
	"net/http"
	"net/url"
	"time"
)

const ServerAddrHttpEh = "https://ieforex.com/dmarket/dsymbol"
const serverAddrPathKline = "/getHistoryData"

// barEHMap EH kline时间粒度
var barEHMap = map[string]string{
	"1m":  "1",
	"5m":  "5",
	"30m": "30",
	"60m": "60",
	"4h":  "240",
	"1D":  "1D",
}

// Post 发送post请求
func Post(instId, bar string) ([]byte, error) {
	path := ServerAddrHttpEh + serverAddrPathKline

	barTime := barEHMap[bar]
	// 获取当前的整点时间搓
	nowTime := time.Now()
	//hourTime := nowTime.Add(-time.Duration(nowTime.Hour()) * time.Hour).Unix()

	// 计算一年前的时间
	oneYearAgo := nowTime.AddDate(-1, 0, 0)

	// 拼接参数
	ehParams := dto.EhParams{
		Id:   "trade." + bar + "." + instId,
		Cmd:  "req",
		Args: []interface{}{"candle." + barTime + "." + instId, 300, oneYearAgo.Unix(), nowTime.Unix()},
	}
	param := url.Values{"message": {utils.ObjToString(ehParams)}}
	result, err := http.PostForm(path, param)
	//formDataBytes := []byte(param.Encode())
	//result, err := http.Post(path, "multipart/form-data", bytes.NewBuffer(formDataBytes))
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return nil, err
	}
	defer result.Body.Close()

	// 读取body数据
	body, err := io.ReadAll(result.Body)
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return nil, err
	}

	// 判断是否状态200
	//if result.StatusCode != 200 {
	//	err = errors.New(fmt.Sprintln("code = ", result.StatusCode))
	//	logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	//	return nil, err
	//}

	//	接口是否正常
	respData := new(dto.RspData)
	err = json.Unmarshal(body, &respData)
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return nil, err
	}

	if respData.Code > http.StatusOK {
		err = errors.New(respData.Msg)
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return nil, err
	}

	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("data", respData))

	return nil, nil
}
