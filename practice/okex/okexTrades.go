package okex

import "encoding/json"

type OkexTrade struct {
	InstId  string `json:"instId"`  //	产品ID
	TradeId string `json:"tradeId"` //	成交ID
	Side    string `json:"side"`    //	成交方向
	Sz      string `json:"sz"`      //	成交数量
	Px      string `json:"px"`      //	成交价格
	Ts      string `json:"ts"`      //	成交时间
}

// Trades 交易量
func (_OkexStruct *OkexStruct) Trades(instId string) ([]*OkexTrade, error) {
	params := map[string]interface{}{"instId": instId}
	resp, err := _OkexStruct.Get("/api/v5/market/trades", params)
	if err != nil {
		return nil, err
	}

	data := make([]*OkexTrade, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}
