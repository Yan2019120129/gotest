package main

type Response struct {
	CacheTime      string   `json:"cacheTime"`
	Agr            string   `json:"agr"`
	Span           string   `json:"span"`
	AllowIntraday  bool     `json:"allow_intraday"`
	AllowInterval  string   `json:"allow_interval"`
	IsIntraday     bool     `json:"isIntraday"`
	ChartFrequency string   `json:"chartFrequency"`
	Series         []Series `json:"series"`
}

type Series struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Shortname            string  `json:"shortname"`
	FullName             string  `json:"full_name"`
	Data                 [][]any `json:"data"` // using any to support different types in the nested arrays
	Unit                 string  `json:"unit"`
	Decimals             int     `json:"decimals"`
	Frequency            string  `json:"frequency"`
	Type                 string  `json:"type"`
	AllowedCandles       bool    `json:"allowed_candles"`
	SupportedResolutions any     `json:"supported_resolutions"`
	Timezone             any     `json:"timezone"`
	HasDaily             bool    `json:"has_daily"`
	AllowedInterval      any     `json:"allowed_interval"`
	Value                any     `json:"value"`
	ConvertedValue       any     `json:"converted_value"`
	Last                 any     `json:"last"`
	Ticker               any     `json:"ticker"`
	Description          any     `json:"description"`
	HasWeeklyAndMonthly  bool    `json:"has_weekly_and_monthly"`
	HasNoVolume          bool    `json:"has_no_volume"`
	HasIntraday          bool    `json:"has_intraday"`
	Industry             any     `json:"industry"`
	Minmov               any     `json:"minmov"`
	Sector               any     `json:"sector"`
	Pricescale           float64 `json:"pricescale"`
	Minmov2              float64 `json:"minmov2"`
	IeconomicsUrl        any     `json:"ieconomics_url"`
	ChartType            any     `json:"chart_type"`
}

// WsResponse 定义握手响应的结构体
type WsResponse struct {
	SID          string   `json:"sid"`
	Upgrades     []string `json:"upgrades"`
	PingInterval int      `json:"pingInterval"`
	PingTimeout  int      `json:"pingTimeout"`
	MaxPayload   int      `json:"maxPayload"`
}

type Params struct {
	Key       string `json:"key"`
	Url       string `json:"url"`
	EIO       string `json:"EIO"`
	Transport string `json:"transport"`
	T         string `json:"t"`
}

type KlineAttrs struct {
	OpenPrice  float64 `json:"openPrice"`  //开盘价格
	HighPrice  float64 `json:"highPrice"`  //最高价格
	LowsPrice  float64 `json:"lowsPrice"`  //最低价格
	ClosePrice float64 `json:"closePrice"` //收盘价格
	Vol        float64 `json:"vol"`        //交易量
	Amount     float64 `json:"amount"`     //成交额
	CreatedAt  int64   `json:"createdAt"`  //开盘时间
}

type CommodityData struct {
	S      string  `json:"s"`      // Symbol (e.g., XAUUSD:CUR)
	P      float64 `json:"p"`      // Price (e.g., 2639.67)
	Nch    float64 `json:"nch"`    // Change in value (e.g., -11.95)
	Pch    float64 `json:"pch"`    // Percentage change (e.g., -0.45)
	Dt     int64   `json:"dt"`     // Date timestamp (e.g., 1734445552026)
	Odt    int64   `json:"odt"`    // Open date timestamp (e.g., 1734445552020)
	Type   string  `json:"type"`   // Type of commodity (e.g., commodity)
	State  string  `json:"state"`  // State (e.g., open)
	Dstate string  `json:"dstate"` // Detail state (e.g., open)
}
