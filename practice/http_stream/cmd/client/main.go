package main

import (
	"context"
	"flag"
	"fmt"
	"gotest/practice/http_stream"
	"log"
	"time"
)

func main() {
	url := flag.String("url", "http://localhost:8080/stream?count=10&interval=500ms", "stream url")
	timeout := flag.Duration("timeout", 30*time.Second, "request timeout")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	client := &http_stream.Client{}
	if err := client.Stream(ctx, *url, func(msg http_stream.Message) error {
		fmt.Printf("[%02d] %s done=%v time=%s\n", msg.Index, msg.Text, msg.Done, msg.Time.Format(time.RFC3339Nano))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
