package okx

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

type Channel string

const (
	// ChannelTicker 行情频道
	ChannelTicker Channel = "tickers"

	// ChannelBooks 深度频道
	ChannelBooks Channel = "books"

	// ChannelBooks5 深度频道
	ChannelBooks5 Channel = "books5"

	// ChannelTrades 交易频道
	ChannelTrades Channel = "trades"

	ChannelKline             Channel = "candle"
	ChannelKlineCandle3M     Channel = "candle3M"
	ChannelKlineCandle1M     Channel = "candle1M"
	ChannelKlineCandle1W     Channel = "candle1W"
	ChannelKlineCandle1D     Channel = "candle1D"
	ChannelKlineCandle2D     Channel = "candle2D"
	ChannelKlineCandle3D     Channel = "candle3D"
	ChannelKlineCandle5D     Channel = "candle5D"
	ChannelKlineCandle12H    Channel = "candle12H"
	ChannelKlineCandle6H     Channel = "candle6H"
	ChannelKlineCandle4H     Channel = "candle4H"
	ChannelKlineCandle2H     Channel = "candle2H"
	ChannelKlineCandle1H     Channel = "candle1H"
	ChannelKlineCandle30m    Channel = "candle30m"
	ChannelKlineCandle15m    Channel = "candle15m"
	ChannelKlineCandle5m     Channel = "candle5m"
	ChannelKlineCandle3m     Channel = "candle3m"
	ChannelKlineCandle1m     Channel = "candle1m"
	ChannelKlineCandle1s     Channel = "candle1s"
	ChannelKlineCandle3Mutc  Channel = "candle3Mutc"
	ChannelKlineCandle1Mutc  Channel = "candle1Mutc"
	ChannelKlineCandle1Wutc  Channel = "candle1Wutc"
	ChannelKlineCandle1Dutc  Channel = "candle1Dutc"
	ChannelKlineCandle2Dutc  Channel = "candle2Dutc"
	ChannelKlineCandle3Dutc  Channel = "candle3Dutc"
	ChannelKlineCandle5Dutc  Channel = "candle5Dutc"
	ChannelKlineCandle12Hutc Channel = "candle12Hutc"
	ChannelKlineCandle6Hutc  Channel = "candle6Hutc"
)

var ChannelMap = map[Channel]string{
	// ChannelTicker 行情频道
	ChannelTicker: "tickers",

	// ChannelBooks 深度频道
	ChannelBooks: "books",

	// ChannelBooks5 深度频道
	ChannelBooks5: "books5",

	// ChannelTrades 交易频道
	ChannelTrades: "trades",
}

var ChannelKlineMap = map[Channel]string{
	ChannelKlineCandle3M:     "candle3M",
	ChannelKlineCandle1M:     "candle1M",
	ChannelKlineCandle1W:     "candle1W",
	ChannelKlineCandle1D:     "candle1D",
	ChannelKlineCandle2D:     "candle2D",
	ChannelKlineCandle3D:     "candle3D",
	ChannelKlineCandle5D:     "candle5D",
	ChannelKlineCandle12H:    "candle12H",
	ChannelKlineCandle6H:     "candle6H",
	ChannelKlineCandle4H:     "candle4H",
	ChannelKlineCandle2H:     "candle2H",
	ChannelKlineCandle1H:     "candle1H",
	ChannelKlineCandle30m:    "candle30m",
	ChannelKlineCandle15m:    "candle15m",
	ChannelKlineCandle5m:     "candle5m",
	ChannelKlineCandle3m:     "candle3m",
	ChannelKlineCandle1m:     "candle1m",
	ChannelKlineCandle1s:     "candle1s",
	ChannelKlineCandle3Mutc:  "candle3Mutc",
	ChannelKlineCandle1Mutc:  "candle1Mutc",
	ChannelKlineCandle1Wutc:  "candle1Wutc",
	ChannelKlineCandle1Dutc:  "candle1Dutc",
	ChannelKlineCandle2Dutc:  "candle2Dutc",
	ChannelKlineCandle3Dutc:  "candle3Dutc",
	ChannelKlineCandle5Dutc:  "candle5Dutc",
	ChannelKlineCandle12Hutc: "candle12Hutc",
	ChannelKlineCandle6Hutc:  "candle6Hutc",
}
