package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/client"
	"strconv"
	"time"
)

var data = []byte(`[{"Name":"tickers","Data":"eyJvcCI6InN1YnNjcmliZSIsImFyZ3MiOlt7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiQlRDLVVTRFQifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiRVRILVVTRFQifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiU09MLVVTRFQifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiVFJYLVVTRFQifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiWFJQLVVTRFQifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiQURBLVVTRFQifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiQVRPTS1VU0RUIn0seyJjaGFubmVsIjoidGlja2VycyIsImluc3RJZCI6IkRPR0UtVVNEVCJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJFVEMtVVNEVCJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJMVEMtVVNEVCJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJVTkktVVNEVCJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJCVEMtVVNEQyJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJFVEgtVVNEQyJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJTT0wtVVNEQyJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJYUlAtVVNEQyJ9LHsiY2hhbm5lbCI6InRpY2tlcnMiLCJpbnN0SWQiOiJET0dFLVVTREMifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiRVRDLVVTREMifSx7ImNoYW5uZWwiOiJ0aWNrZXJzIiwiaW5zdElkIjoiTFRDLVVTREMifV19"},{"Name":"books","Data":"eyJvcCI6InN1YnNjcmliZSIsImFyZ3MiOlt7ImNoYW5uZWwiOiJib29rcyIsImluc3RJZCI6IkJUQy1VU0RUIn0seyJjaGFubmVsIjoiYm9va3MiLCJpbnN0SWQiOiJFVEgtVVNEVCJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiU09MLVVTRFQifSx7ImNoYW5uZWwiOiJib29rcyIsImluc3RJZCI6IlRSWC1VU0RUIn0seyJjaGFubmVsIjoiYm9va3MiLCJpbnN0SWQiOiJYUlAtVVNEVCJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiQURBLVVTRFQifSx7ImNoYW5uZWwiOiJib29rcyIsImluc3RJZCI6IkFUT00tVVNEVCJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiRE9HRS1VU0RUIn0seyJjaGFubmVsIjoiYm9va3MiLCJpbnN0SWQiOiJFVEMtVVNEVCJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiTFRDLVVTRFQifSx7ImNoYW5uZWwiOiJib29rcyIsImluc3RJZCI6IlVOSS1VU0RUIn0seyJjaGFubmVsIjoiYm9va3MiLCJpbnN0SWQiOiJCVEMtVVNEQyJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiRVRILVVTREMifSx7ImNoYW5uZWwiOiJib29rcyIsImluc3RJZCI6IlNPTC1VU0RDIn0seyJjaGFubmVsIjoiYm9va3MiLCJpbnN0SWQiOiJYUlAtVVNEQyJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiRE9HRS1VU0RDIn0seyJjaGFubmVsIjoiYm9va3MiLCJpbnN0SWQiOiJFVEMtVVNEQyJ9LHsiY2hhbm5lbCI6ImJvb2tzIiwiaW5zdElkIjoiTFRDLVVTREMifV19"},{"Name":"trades","Data":"eyJvcCI6InN1YnNjcmliZSIsImFyZ3MiOlt7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJCVEMtVVNEVCJ9LHsiY2hhbm5lbCI6InRyYWRlcyIsImluc3RJZCI6IkVUSC1VU0RUIn0seyJjaGFubmVsIjoidHJhZGVzIiwiaW5zdElkIjoiU09MLVVTRFQifSx7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJUUlgtVVNEVCJ9LHsiY2hhbm5lbCI6InRyYWRlcyIsImluc3RJZCI6IlhSUC1VU0RUIn0seyJjaGFubmVsIjoidHJhZGVzIiwiaW5zdElkIjoiQURBLVVTRFQifSx7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJBVE9NLVVTRFQifSx7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJET0dFLVVTRFQifSx7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJFVEMtVVNEVCJ9LHsiY2hhbm5lbCI6InRyYWRlcyIsImluc3RJZCI6IkxUQy1VU0RUIn0seyJjaGFubmVsIjoidHJhZGVzIiwiaW5zdElkIjoiVU5JLVVTRFQifSx7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJCVEMtVVNEQyJ9LHsiY2hhbm5lbCI6InRyYWRlcyIsImluc3RJZCI6IkVUSC1VU0RDIn0seyJjaGFubmVsIjoidHJhZGVzIiwiaW5zdElkIjoiU09MLVVTREMifSx7ImNoYW5uZWwiOiJ0cmFkZXMiLCJpbnN0SWQiOiJYUlAtVVNEQyJ9LHsiY2hhbm5lbCI6InRyYWRlcyIsImluc3RJZCI6IkRPR0UtVVNEQyJ9LHsiY2hhbm5lbCI6InRyYWRlcyIsImluc3RJZCI6IkVUQy1VU0RDIn0seyJjaGFubmVsIjoidHJhZGVzIiwiaW5zdElkIjoiTFRDLVVTREMifV19"}]`)

// SubscribeArg 币种订阅频道。
type SubscribeArg struct {
	Channel string `json:"channel"` // 订阅的通道
	InstID  string `json:"instId"`  // 货币类型
}

// SubscribeRespJson 订阅返回数据
type SubscribeRespJson struct {
	Arg  *SubscribeArg   `json:"arg"`
	Data json.RawMessage `json:"data"`
}

// OkxTickers 产品行情信息
type OkxTickers struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        string `json:"ts"`
}

func main() {
	symbolList := make([]*client.Subscribe, 0)
	_ = json.Unmarshal(data, &symbolList)
	client.NewSocketClient("wss://ws.okx.com:8443/ws/v5/public").InitSubscribes(symbolList).SetWebSocketMessageFunc(func(rdsConn redis.Conn, bytes []byte) error {
		//	解析数据
		respJson := &SubscribeRespJson{}
		_ = json.Unmarshal(bytes, &respJson)
		switch respJson.Arg.Channel {
		case "tickers":
			okExchangeTickers := make([]*OkxTickers, 0)
			_ = json.Unmarshal(respJson.Data, &okExchangeTickers)
			if len(okExchangeTickers) > 0 {
				symbol := okExchangeTickers[0].InstId
				ts, _ := strconv.Atoi(okExchangeTickers[0].Ts)
				newTime := time.UnixMilli(int64(ts))
				if symbol == "BTC-USDT" {
					fmt.Println(symbol, newTime.Format(time.DateTime))
				}
			}
		}
		return nil
	}).Connect()
	time.Sleep(5 * time.Hour)
}
