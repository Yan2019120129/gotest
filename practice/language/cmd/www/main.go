package main

import (
	"github.com/gorilla/mux"
	// 引入internal/translations包，确保init()函数被调用
	_ "language/internal/translations"

	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{locale}", handleHome)

	// 使用路由实例启动HTTP服务器。
	log.Println("starting server on :4018...")
	if err := http.ListenAndServe(":4018", r); err != nil {
		log.Fatal(err)
	}
}
