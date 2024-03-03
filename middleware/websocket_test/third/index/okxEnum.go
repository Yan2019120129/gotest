package index

import (
	"maps"
)

// ServerOkxAddr 产品行情地址。
const (
	// ServerOkxAddr okx 行情websocket 地址
	ServerOkxAddr = "wss://ws.okx.com:8443/ws/v5/public"

	// ServerCandleAndTradeAddr okx 行业websocket 地址
	ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"

	// OpSubscribe 订阅
	OpSubscribe = "subscribe"

	// OpUnsubscribe 取消订阅
	OpUnsubscribe = "unsubscribe"
)

const (
	// ChannelTicker 行情频道
	ChannelTicker = "tickers"

	// ChannelBooks 深度频道
	ChannelBooks = "books"

	// ChannelBooks5 深度频道
	ChannelBooks5 = "books5"

	// ChannelTrades 交易频道
	ChannelTrades = "trades"

	ChannelKline             = "candle"
	ChannelKlineCandle3M     = "candle3M"
	ChannelKlineCandle1M     = "candle1M"
	ChannelKlineCandle1W     = "candle1W"
	ChannelKlineCandle1D     = "candle1D"
	ChannelKlineCandle2D     = "candle2D"
	ChannelKlineCandle3D     = "candle3D"
	ChannelKlineCandle5D     = "candle5D"
	ChannelKlineCandle12H    = "candle12H"
	ChannelKlineCandle6H     = "candle6H"
	ChannelKlineCandle4H     = "candle4H"
	ChannelKlineCandle2H     = "candle2H"
	ChannelKlineCandle1H     = "candle1H"
	ChannelKlineCandle30m    = "candle30m"
	ChannelKlineCandle15m    = "candle15m"
	ChannelKlineCandle5m     = "candle5m"
	ChannelKlineCandle3m     = "candle3m"
	ChannelKlineCandle1m     = "candle1m"
	ChannelKlineCandle1s     = "candle1s"
	ChannelKlineCandle3Mutc  = "candle3Mutc"
	ChannelKlineCandle1Mutc  = "candle1Mutc"
	ChannelKlineCandle1Wutc  = "candle1Wutc"
	ChannelKlineCandle1Dutc  = "candle1Dutc"
	ChannelKlineCandle2Dutc  = "candle2Dutc"
	ChannelKlineCandle3Dutc  = "candle3Dutc"
	ChannelKlineCandle5Dutc  = "candle5Dutc"
	ChannelKlineCandle12Hutc = "candle12Hutc"
	ChannelKlineCandle6Hutc  = "candle6Hutc"
)

// ChannelMap 行情通道
var ChannelMap = map[string]string{
	// ChannelTicker 行情频道
	ChannelTicker: "tickers",

	// ChannelBooks 深度频道
	ChannelBooks: "books",

	// ChannelBooks5 深度频道
	ChannelBooks5: "books5",

	// ChannelTrades 交易频道
	ChannelTrades: "trades",
}

// ChannelKlineMap k线图通道
var ChannelKlineMap = map[string]string{
	ChannelKlineCandle1m:  "candle1m",
	ChannelKlineCandle5m:  "candle5m",
	ChannelKlineCandle30m: "candle30m",
	ChannelKlineCandle1H:  "candle1H",
	ChannelKlineCandle4H:  "candle4H",
	ChannelKlineCandle15m: "candle1D",
}

// ChannelAllMap 全部的通道
var ChannelAllMap = marge()

func marge() map[string]string {
	tempMap := make(map[string]string)
	maps.Copy(tempMap, ChannelMap)
	maps.Copy(tempMap, ChannelKlineMap)
	return tempMap
}
