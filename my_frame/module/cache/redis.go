package cache

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/my_frame/config"
	"sync"
	"time"
)

// 初始化redis连接池,保证全局单例。
var _once sync.Once

// RdsPool redis连接池单例,保证全局单例,保证全局单例,保证全局单例。
var RdsPool *redis.Pool

// RdsPubSubConn redis连接池单例,保证全局单例,保证全局单例,保证全局单例。
var RdsPubSubConn *redis.PubSubConn

// 获取redis配置
var cfg = config.GetRedisConfig()

// Init 初始化Redis。
func init() {
	if RdsPool == nil {
		_once.Do(func() {
			RdsPool = &redis.Pool{
				MaxIdle:     cfg.Poll.MaxIdleConn,
				MaxActive:   cfg.Poll.MaxOpenConn,
				IdleTimeout: time.Duration(cfg.Poll.ConnectTimeout) * time.Second,
				Wait:        false,
				Dial: func() (redis.Conn, error) {
					if conn, err := redis.Dial(
						cfg.Poll.Network,
						getDsn(&cfg.Poll),
						redis.DialPassword(cfg.Poll.Pass),
						redis.DialDatabase(cfg.Poll.DbName),
						redis.DialConnectTimeout(time.Duration(cfg.Poll.ConnectTimeout)*time.Second),
						redis.DialReadTimeout(time.Duration(cfg.Poll.ReadTimeout)*time.Second),
						redis.DialWriteTimeout(time.Duration(cfg.Poll.WriteTimeout)*time.Second),
					); err == nil {
						return conn, nil
					}
					return nil, errors.New("redis初始化失败！！！")
				},
			}

			// 是否启用订阅
			if cfg.UsePub {
				RdsPubSubConn = &redis.PubSubConn{Conn: RdsPool.Get()}
				fmt.Printf("内存地址：%p----->Rds开启订阅！！！\n", RdsPubSubConn)
			}
			fmt.Printf("内存地址：%p----->Rds连接池初始化成功！！！\n", RdsPool)
		})
	} else {
		fmt.Println("已经存在Res实例！！！")
	}
}

// getDsn 获取dsn字符串。
func getDsn(cfg *config.RedisPollConfig) string {
	return fmt.Sprintf("%v:%v", cfg.Network, cfg.Port)
}
