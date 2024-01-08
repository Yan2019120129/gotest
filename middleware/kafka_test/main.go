package main

import (
	"gotest/middleware/kafka_test/exaple"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go exaple.Producer()
	wg.Add(1)
	wg.Wait()
}
