package main

import (
	"fmt"
	"github.com/gorilla/mux"
	main2 "gotest/websocket/one"
	"net/http"
)

// TestWebsocket 测试websocket
func main() {
	router := mux.NewRouter()
	go main2.h.run()
	router.HandleFunc("/ws", main2.myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
