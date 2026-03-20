package utils

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Ws struct {
	header http.Header
	dialer *websocket.Dialer
	conn   *websocket.Conn
	url    string
	params url.Values
}

func NewWs(u string) *Ws {
	return &Ws{
		url:    u,
		dialer: websocket.DefaultDialer,
	}
}

func (i *Ws) Run() *Ws {
	if i.params != nil {
		i.url += "?" + i.params.Encode()
	}

	var err error
	i.conn, _, err = i.dialer.Dial(i.url, i.header)
	if err != nil {
		panic(err)
	}
	return i
}

// AddParam 添加请求参数
func (i *Ws) AddParam(key string, val string) *Ws {
	if i.params == nil {
		i.params = make(url.Values)
	}
	i.params.Set(key, val)
	return i
}

// Set 设置请求头信息
func (i *Ws) Set(key, val string) *Ws {
	if i.header == nil {
		i.header = make(http.Header)
	}
	i.header.Set(key, val)
	return i
}

func (i *Ws) Send(s string) *Ws {
	err := i.conn.WriteMessage(websocket.TextMessage, []byte(s))
	if err != nil {
		return i
	}
	return i
}

func (i *Ws) Ping(t time.Duration, s string) *Ws {
	go func() {
		for {
			err := i.conn.WriteMessage(websocket.PingMessage, []byte(s))
			if err != nil {
				_ = i.conn.Close()
				return
			}
			time.Sleep(t * time.Second)
		}
	}()
	return i
}

func (i *Ws) Read(fu func([]byte)) {
	for {
		_, message, err := i.conn.ReadMessage()
		if err != nil {
			_ = i.conn.Close()
			return
		}
		fu(message)
	}
}
