package okex

import (
	"encoding/json"
	"strconv"
)

type KlineAttrs struct {
	OpenPrice  float64 `json:"openPrice"`  //开盘价格
	HighPrice  float64 `json:"highPrice"`  //最高价格
	LowsPrice  float64 `json:"lowsPrice"`  //最低价格
	ClosePrice float64 `json:"closePrice"` //收盘价格
	Vol        float64 `json:"vol"`        //交易量
	Amount     float64 `json:"amount"`     //成交额
	CreatedAt  int64   `json:"createdAt"`  //开盘时间
}

// Kline k线数据
func (_OkexStruct *OkexStruct) Kline(instId string, bar string) ([]*KlineAttrs, error) {
	params := map[string]interface{}{"instId": instId, "limit": "300", "bar": bar}
	resp, err := _OkexStruct.Get("/api/v5/market/candles", params)
	if err != nil {
		return nil, err
	}

	oldData := make([][]string, 0)
	_ = json.Unmarshal(resp, &oldData)

	data := make([]*KlineAttrs, 0)
	for i := 0; i < len(oldData); i++ {
		kline := oldData[i]

		openPrice, _ := strconv.ParseFloat(kline[1], 64)
		highPrice, _ := strconv.ParseFloat(kline[2], 64)
		lowsPrice, _ := strconv.ParseFloat(kline[3], 64)
		closePrice, _ := strconv.ParseFloat(kline[4], 64)
		vol, _ := strconv.ParseFloat(kline[6], 64)
		quote, _ := strconv.ParseFloat(kline[7], 64)

		createTime, _ := strconv.ParseInt(kline[0], 10, 64)
		data = append(data, &KlineAttrs{
			OpenPrice:  openPrice,
			HighPrice:  highPrice,
			LowsPrice:  lowsPrice,
			ClosePrice: closePrice,
			Vol:        vol,
			Amount:     quote,
			CreatedAt:  createTime / 1000,
		})
	}

	return data, nil
}
