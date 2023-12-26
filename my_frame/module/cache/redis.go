package cache

import (
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

// Init 初始化Redis。
func init() {
	if RdsPool == nil {
		_once.Do(func() {
			var cfg = config.GetRedisConfig()
			RdsPool = &redis.Pool{
				MaxIdle:     cfg.Pool.MaxIdleConn,
				MaxActive:   cfg.Pool.MaxOpenConn,
				IdleTimeout: time.Duration(cfg.Pool.ConnectTimeout) * time.Second,
				Wait:        false,
				Dial: func() (redis.Conn, error) {
					if conn, err := redis.Dial(
						cfg.Pool.Network,
						getDsn(&cfg.Pool),
						redis.DialPassword(cfg.Pool.Pass),
						redis.DialDatabase(cfg.Pool.DbName),
						redis.DialConnectTimeout(time.Duration(cfg.Pool.ConnectTimeout)*time.Second),
						redis.DialReadTimeout(time.Duration(cfg.Pool.ReadTimeout)*time.Second),
						redis.DialWriteTimeout(time.Duration(cfg.Pool.WriteTimeout)*time.Second),
					); err == nil {
						return conn, nil
					} else {
						return nil, err
					}
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
func getDsn(cfg *config.RedisPoolConfig) string {
	return fmt.Sprintf("%v:%v", cfg.Network, cfg.Port)
}
