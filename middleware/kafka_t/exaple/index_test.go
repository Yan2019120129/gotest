package exaple

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

// TestTopic 生产者
func TestTopic(t *testing.T) {
	go Topic()
	wg.Add(1)
	wg.Wait()
}

// TestProducer 生产者
func TestProducer(t *testing.T) {
	go Producer()
	wg.Add(1)
	wg.Wait()
}

// TestConsumer 消费者
func TestConsumer(t *testing.T) {
	go Consumer()
	wg.Add(1)
	wg.Wait()
}
