package indexserver

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/frame/my_frame/app/admin/service/dto"
	"gotest/frame/my_frame/module/cache"
)

// SubRds 订阅信息
func SubRds(params *dto.RdsPublishAndParams) (interface{}, error) {
	rdsConn := cache.RdsPubSubConn
	defer rdsConn.Close()
	ctx := context.Background()
	if err := rdsConn.Subscribe(ctx, params.Channel); err != nil {
		return nil, err
	}
	go func() {
		for {
			switch v := rdsConn.Receive().(type) {
			case redis.Message:
				fmt.Println("data:", v.Channel, v.Data)
			case redis.Subscription:
				fmt.Println("订阅:", v.Count)
			}
		}
	}()
	return "ok", nil
}

// Publish 发布信息
func Publish(params *dto.RdsPublishAndParams) (interface{}, error) {
	rdsConn := cache.RdsPubSubConn
	defer rdsConn.Close()
	ctx := context.Background()
	if err := rdsConn.PUnsubscribe(ctx, params.Channel, params.Message); err != nil {
		return nil, err
	}

	return "ok", nil
}
