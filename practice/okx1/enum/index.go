package enum

const (
	OriginWsUrlOKX  = "wss://ws.okx.com:8443"
	TickersWsUrlOKX = OriginWsUrlOKX + "/ws/v5/public"
	TradesWsUrlOKX  = OriginWsUrlOKX + "/ws/v5/public"
	KlineWsUrlOKX   = OriginWsUrlOKX + "/ws/v5/business"

	OKXChannelTicker   = "tickers"
	OKXChannelTrades   = "trades"
	OKXChannelCandle3M = "trades"
)
