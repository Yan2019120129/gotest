package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/common/config"
	"log"
	"sync"
	"time"
)

// 初始化redis连接池,保证全局单例。
var _once sync.Once

// RdsPool redis连接池单例,保证全局单例,保证全局单例,保证全局单例。
var RdsPool *redis.Pool

// RdsPubSubConn redis连接池单例,保证全局单例,保证全局单例,保证全局单例。
var RdsPubSubConn *redis.PubSubConn

type instance struct {
	RdsPubSubConn *redis.PubSubConn // redis 实例
	RdsPool       *redis.Pool
}

var Instance = &instance{}

// Init 初始化Redis。
func init() {
	if RdsPool == nil {
		_once.Do(func() {
			var cfg = config.GetRedis()
			RdsPool = &redis.Pool{
				MaxIdle:     cfg.Pool.MaxIdleConn,
				MaxActive:   cfg.Pool.MaxOpenConn,
				IdleTimeout: time.Duration(cfg.Pool.ConnectTimeout) * time.Second,
				Wait:        false,
				Dial: func() (redis.Conn, error) {
					host := fmt.Sprintf("%v:%v", cfg.Pool.Server, cfg.Pool.Port)
					conn, err := redis.Dial(
						cfg.Pool.Network,
						host,
						redis.DialPassword(cfg.Pool.Pass),
						redis.DialDatabase(cfg.Pool.DbName),
						redis.DialConnectTimeout(time.Duration(cfg.Pool.ConnectTimeout)*time.Second),
						redis.DialReadTimeout(time.Duration(cfg.Pool.ReadTimeout)*time.Second),
						redis.DialWriteTimeout(time.Duration(cfg.Pool.WriteTimeout)*time.Second),
					)
					if err != nil {
						panic(err)
						return nil, err
					}
					return conn, nil
				},
			}
			Instance.RdsPool = RdsPool
			// 是否启用订阅
			if cfg.UsePub {
				RdsPubSubConn = &redis.PubSubConn{Conn: RdsPool.Get()}
				Instance.RdsPubSubConn = RdsPubSubConn

				fmt.Printf("内存地址：%p----->Rds开启订阅！！！\n", RdsPubSubConn)
			}
			fmt.Printf("内存地址：%p----->Rds连接池初始化成功！！！\n", RdsPool)
		})
	} else {
		fmt.Println("已经存在Res实例！！！")
	}
}

// Subscribe 订阅消息
func (i *instance) Subscribe(fun func(message []byte, channel string), channels ...interface{}) error {
	if err := i.RdsPubSubConn.Subscribe(redis.Args{}.Add(channels...)...); err != nil {
		return err
	}
	for {
		switch v := i.RdsPubSubConn.Receive().(type) {
		case redis.Message:
			fun(v.Data, v.Channel)
		case redis.Subscription:
			fmt.Printf("订阅频道: %s，订阅数量: %d\n", v.Channel, v.Count)
		case error:
			fmt.Println("接收消息错误:", v)
			return v
		}
	}
}

// UnSubscribe 取消订阅
func (i *instance) UnSubscribe(channels ...interface{}) error {
	if err := i.RdsPubSubConn.Unsubscribe(redis.Args{}.Add(channels...)...); err != nil {
		return err
	}
	return nil
}

// Publish 发布信息
func (i *instance) Publish(message []byte, channels ...interface{}) {
	rdsConn := i.RdsPool.Get()
	defer rdsConn.Close()
	for _, v := range channels {
		if err := rdsConn.Send("PUBLISH", v, message); err != nil {
			log.Println(err)
			return
		}
	}
}
