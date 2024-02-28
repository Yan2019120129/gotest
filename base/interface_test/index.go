package interface_test

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var message = []byte(`{
	"op": "subscribe",
	"data": [
		[
            "1708617300000",
            "51573.7",
            "51621.8",
            "51573.7",
            "51602.5",
            "32.19050846",
            "1661243.0761018",
            "1661243.0761018",
            "1"
        ],
        [
            "1708617000000",
            "51531.2",
            "51576.9",
            "51506.1",
            "51573.6",
            "37.36719661",
            "1926220.73216047",
            "1926220.73216047",
            "1"
        ]
	]
}`)

// Message 消息体
type Message struct {
	Op   string        `json:"op"`   //	方法名称
	Data []interface{} `json:"data"` //	方法参数
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

// InterfaceToObj interface转换结构体类型
func InterfaceToObj(data interface{}, obj interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataBytes, obj)
	return err
}

// ToKline 二维字符串数组转换为KlineData
func ToKline(kline []string) (data *KlineData) {
	createTime, err := strconv.ParseInt(kline[0], 10, 64)
	if err != nil {
		panic(err)
	}

	openPrice, err := strconv.ParseFloat(kline[1], 64)
	if err != nil {
		panic(err)
	}

	highPrice, err := strconv.ParseFloat(kline[2], 64)
	if err != nil {
		panic(err)
	}

	lowsPrice, err := strconv.ParseFloat(kline[3], 64)
	if err != nil {
		panic(err)
	}

	closePrice, err := strconv.ParseFloat(kline[4], 64)
	if err != nil {
		panic(err)
	}

	vol, err := strconv.ParseFloat(kline[6], 64)
	if err != nil {
		panic(err)
	}

	quote, err := strconv.ParseFloat(kline[7], 64)
	if err != nil {
		panic(err)
	}
	return &KlineData{
		OpenPrice:  openPrice,
		HighPrice:  highPrice,
		LowsPrice:  lowsPrice,
		ClosePrice: closePrice,
		Vol:        vol,
		Amount:     quote,
		CreatedAt:  createTime / 1000,
	}
}

// InterfaceToStruct 接口转换为结构体
func InterfaceToStruct() {
	data := Message{}
	if err := json.Unmarshal(message, &data); err != nil {
		panic(err)
	}

	tempData := []string{}
	if err := InterfaceToObj(data.Data[0], &tempData); err != nil {
		panic(err)
	}

	fmt.Println("data", tempData)

}
