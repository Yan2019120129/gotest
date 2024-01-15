package okex

import "encoding/json"

type OkexBook struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

// Trades 交易深度
func (_OkexStruct *OkexStruct) Books(instId string) ([]*OkexBook, error) {
	params := map[string]interface{}{"instId": instId, "sz": "60"}
	resp, err := _OkexStruct.Get("/api/v5/market/books", params)
	if err != nil {
		return nil, err
	}

	data := make([]*OkexBook, 0)
	_ = json.Unmarshal(resp, &data)
	return data, nil
}
