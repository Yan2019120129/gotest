package main

import (
	"flag"
	"fmt"
	"gotest/practice/http_stream"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	flag.Parse()

	fmt.Printf("stream server listening on http://localhost%s/stream\n", *addr)
	if err := http.ListenAndServe(*addr, http_stream.NewServeMux()); err != nil {
		log.Fatal(err)
	}
}
