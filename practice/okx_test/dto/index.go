package dto

type TickerParams struct {
	Op   string `json:"op"`
	Args []Arg  `json:"args"`
}

type TickerData struct {
	Arg  Arg     `json:"arg"`
	Data []Datum `json:"data"`
}

type Arg struct {
	Channel string `json:"channel"`
	InstID  string `json:"instId"`
}

type Datum struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
}
