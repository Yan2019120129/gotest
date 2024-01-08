package exaple

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
	"gotest/common/module/logger"
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
var topic = "myTopic"

// Topic 逻辑分类
func Topic() {
	// 创建 Kafka AdminClient
	adminClient, err := kafka.NewAdminClient(config)
	if err != nil {
		logger.Logger.Warn(err.Error())
	}

	adminName := adminClient.String()
	logger.Logger.Info("连接成功：" + adminName)

	// 上下文超时设置，超出时间关闭程序
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	// 检查发送信息是否有问题
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logger.Logger.Warn("错误信息：", zap.Error(ev.TopicPartition.Error))
				} else {
					logger.Logger.Info("发送信息：", zap.Reflect("Delivered message to", ev))
				}
			}
		}
	}()

	// 手动输入信息
	for {
		message := ""
		fmt.Print("请输入信息：")
		if _, err = fmt.Scan(&message); err != nil {
			logger.Logger.Warn("错误信息：", zap.Error(err))
		}
		if err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil); err != nil {
			logger.Logger.Warn("错误信息：", zap.Error(err))
			return
		}
	}
	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

// Consumer 消费者
func Consumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": serverAddr,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		logger.Logger.Warn("错误信息：", zap.String("subscribes", err.Error()))
	}
	defer c.Close()

	if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
		logger.Logger.Warn("错误信息：", zap.String("subscribes", err.Error()))
	}

	for {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			logger.Logger.Info("信息：", zap.ByteString("Message", msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			logger.Logger.Debug("错误信息：", zap.String("error", err.Error()), zap.Reflect("error", msg))
			return
		}
	}
}
