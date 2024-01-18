package okex

import (
	"go.uber.org/zap"
	"gotest/common/module/gorm/database"
	"gotest/common/module/log/zap_log"
	"gotest/practice/okex/dto"
	"strings"
)

type Subscribe struct {
	Op   string           `json:"op"`
	Args []*SubscribeArgs `json:"args"`
}

type SubscribeArgs struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type SubscribeData struct {
	Arg    *SubscribeArgs `json:"arg"`
	Action string         `json:"action"`
	Data   interface{}    `json:"data"`
}

// SubscribeTickers 订阅行情频道
func (_OkexStruct *OkexStruct) SubscribeTickers() *OkexStruct {

	productList := make([]*dto.ProductData, 0)
	if result := database.DB.Model(&dto.Product{}).Select("name", "id").Where("type = ?", dto.ProductTypeOkex).Find(&productList); result.Error != nil {
		zap_log.Logger.Warn("warn", zap.Error(result.Error))
	}
	for _, v := range productList {
		v.Name = strings.ReplaceAll(v.Name, "/", "-")
		// 订阅 行情频道
		_OkexStruct.Subscribe(&Subscribe{
			Op: "subscribe",
			Args: []*SubscribeArgs{
				{Channel: "tickers", InstId: v.Name},
			},
		})

		// 订阅交易频道
		_OkexStruct.Subscribe(&Subscribe{
			Op: "subscribe",
			Args: []*SubscribeArgs{
				{Channel: "trades", InstId: v.Name},
			},
		})

		// 订阅深度频道
		_OkexStruct.Subscribe(&Subscribe{
			Op: "subscribe",
			Args: []*SubscribeArgs{
				{Channel: "books", InstId: v.Name},
			},
		})
	}
	return _OkexStruct
}
