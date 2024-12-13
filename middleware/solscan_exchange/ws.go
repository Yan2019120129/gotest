package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type Ws struct {
	conn *websocket.Conn
	url  string
}

func NewWs(u string, proxyUrl string) *Ws {
	var dialer *websocket.Dialer
	//if proxyUrl != "" {
	//	proxy, _ := url.Parse(proxyUrl)
	//	dialer = &websocket.Dialer{Proxy: http.ProxyURL(proxy)}
	//} else {
	dialer = websocket.DefaultDialer
	//}

	conn, _, err := dialer.Dial(u, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Ws{
		conn: conn,
	}
}

func (i *Ws) Send(s string) *Ws {
	err := i.conn.WriteMessage(websocket.TextMessage, []byte(s))
	if err != nil {
		log.Fatal(err)
		return i
	}
	return i
}

func (i *Ws) Read(fu func([]byte)) {
	for {
		_, message, err := i.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			i.conn.Close()
			return
		}
		fu(message)
	}
}
