package json_t

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/goccy/go-json"
	"gotest/middleware/mysql_t/dto"
	"log"
	"time"
)

// TestJson 测试json解析
func TestJson() {
	data := &dto.ProductData{
		InstId:    gofakeit.MonthString(),
		Last:      gofakeit.Float64Range(1000, 10000),
		LastSz:    gofakeit.Float64Range(10, 100),
		Open24h:   gofakeit.Float64Range(500, 1000),
		High24h:   gofakeit.Float64Range(500, 1000),
		Low24h:    gofakeit.Float64Range(500, 1000),
		Vol24h:    gofakeit.Float64Range(500, 1000),
		Amount24h: gofakeit.Float64Range(500, 1000),
		Ts:        time.Now().Unix(),
	}
	//productDataBytes, err := json.Marshal(data)
	//if err != nil {
	//	panic(err)
	//}
	test := map[string]interface{}{
		"id":          1,
		"category_id": 2,
		"images":      []string{"/assets/currency/doge.png"},
		"name":        gofakeit.Name(),
		"money":       gofakeit.Float64Range(10, 100),
		"type":        1,
		"sales":       1,
		"nums":        1,
		"used":        1,
		"total":       1,
		"isCollect":   true,
		"data":        data,
	}
	dataBytes, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	product := dto.IndexData{}
	if err = json.Unmarshal(dataBytes, &product); err != nil {
		panic(err)
	}
	fmt.Println("data:", product)
}

// booksData 返回的推送数据
type booksData struct {
	Arg    arg         `json:"arg"`
	Action string      `json:"action"`
	Data   []BooksData `json:"data"`
}

// BooksData 深度内部数据
type BooksData struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Ts        string     `json:"ts"`
	Checksum  int        `json:"checksum"`
	PrevSeqId int        `json:"prevSeqId"`
	SeqId     int        `json:"seqId"`
}

// Arg 币种订阅频道。
type arg struct {
	Channel string `json:"channel"` // 订阅的通道
	InstID  string `json:"instId"`  // 货币类型
}

// StringToJson 字符串转换json
func StringToJson() {
	data := []byte(`{"arg":{"channel":"books","instId":"BTC-USDT"},"action":"update","data":[{"asks":[["42631.4","0.02790313","0","1"],["42637.7","0","0","0"],["42638.3","0.27038272","0","1"],["42647.7","0.7","0","1"],["42650.8","0.02501619","0","1"],["42759.1","0","0","0"],["42759.3","0","0"]],"bids":[["42607.7","0","0","0"],["42606.5","0.64830961","0","1"],["42582.2","0","0","0"],["42578","0.00361166","0","1"],["42577.9","0.05090297","0","1"],["42577","0","0","0"],["42576.6","1.20878136","0","5"],["42537.3","0.82198593","0","1"],["42507","0","0","0"]],"ts":"1705206812906","checksum":-682828320,"seqId":18330474721,"prevSeqId":18330474702}]}`)
	//fmt.Printf("data:%T\n", data)
	message := booksData{}
	b := BooksData{}
	channel := ""
	for {
		if err := json.Unmarshal(data, &channel); err == nil {

			log.Println("data2", channel)
		}
		if err := json.Unmarshal(data, &message); err == nil {

			log.Println("data3", message.Data)
		}
		if err := json.Unmarshal(data, &b); err == nil {

			log.Println("data4", b)
		}
		log.Println("channel:", channel)
		break
	}
}

// Message 消息体
type Message struct {
	Op   string      `json:"op"`   //	方法名称
	Data interface{} `json:"data"` //	方法参数
}

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	Channel string `json:"channel"` //	订阅通道
	Arg     string `json:"arg"`     //	订阅标识
}

// InterfaceToObj 测试结构体转
func InterfaceToObj() {
	msg := []byte(`{
	"op": "subscribe",
	"data": [
		{
			"channel": "tickers",
			"arg": "LTC-BTC"
		}
	]
}`)

	temp := Message{}
	_ = json.Unmarshal(msg, &temp)

	fmt.Printf("%T", msg)
	fmt.Println("msg", msg)
	fmt.Printf("%T", temp)
	fmt.Println("temp", temp)

	dateTemp, _ := json.Marshal(temp.Data)
	subscribeMessageTemp := make([]SubscribeMessage, 0)
	_ = json.Unmarshal(dateTemp, &subscribeMessageTemp)
	fmt.Printf("%T", subscribeMessageTemp)
	fmt.Println("temp", subscribeMessageTemp)
}
