package kafka_test_test

import (
	"gotest/middleware/kafka_test"
	"testing"
)

// TestTopic 生产者
func TestTopic(t *testing.T) {
	kafka_test.Topic()
}

// TestProducer 生产者
func TestProducer(t *testing.T) {
	kafka_test.Producer()
}

// TestConsumer 消费者
func TestConsumer(t *testing.T) {
	kafka_test.Consumer()
}
