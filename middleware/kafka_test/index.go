package kafka_test

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
	"gotest/module/logger"
	"time"
)

// serverAddr 服务地址
var serverAddr = "192.168.216.139:9092"

//var serverAddr = "47.101.70.217:1018"

// config 基本配置
var config = &kafka.ConfigMap{
	"bootstrap.servers": serverAddr,
}

// 主题
var topic = "my-topic"

// Topic 逻辑分类
func Topic() {
	// 创建 Kafka AdminClient
	adminClient, err := kafka.NewAdminClient(config)
	if err != nil {
		logger.Logger.Warn(err.Error())
	}

	// 上下文超时设置，超出时间关闭程序
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminName := adminClient.String()
	logger.Logger.Info(adminName)

	// 使用 AdminClient 创建 Topic
	results, err := adminClient.CreateTopics(ctx, []kafka.TopicSpecification{
		{
			Topic:             topic, // 要创建的 Topic 名称
			NumPartitions:     3,     // 分区数
			ReplicationFactor: 1,     // 副本数
		},
	})
	if err != nil {
		logger.Logger.Warn(err.Error())
	}

	// 检查创建 Topic 的结果
	for _, result := range results {
		if result.Error.IsTimeout() {
			logger.Logger.Warn(err.Error())
		}
		logger.Logger.Info(result.Topic)
	}

	// 关闭 Kafka AdminClient
	adminClient.Close()
}

// Producer 生产者
func Producer() {
	p, err := kafka.NewProducer(config)
	if err != nil {
		logger.Logger.Warn("错误信息：", zap.Error(err))
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logger.Logger.Warn("错误信息：", zap.Error(ev.TopicPartition.Error))
				} else {
					logger.Logger.Warn("错误信息：", zap.Reflect("Delivered message to", ev))
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		if err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil); err != nil {
			return
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

// Consumer 消费者
func Consumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"api.version.request": true,
		"bootstrap.servers":   serverAddr,
		"group.id":            "myGroup",
		"auto.offset.reset":   "earliest",
		"security.protocol":   "PLAINTEXT",
	})
	if err != nil {
		logger.Logger.Warn("错误信息：", zap.String("subscribes", err.Error()))
	}
	defer c.Close()

	if err = c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil); err != nil {
		logger.Logger.Warn("错误信息：", zap.String("subscribes", err.Error()))
	}

	for {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			logger.Logger.Debug("信息：", zap.Reflect("Message", msg.TopicPartition), zap.ByteString("Message", msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			logger.Logger.Debug("错误信息：", zap.String("error", err.Error()), zap.Reflect("error", msg))
			return
		}
	}
}