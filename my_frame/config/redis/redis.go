package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/my_frame/config"
	"time"
)

var Rds *redis.Pool

// InitRedis 初始化Redis
func InitRedis() {
	cfg := config.Cfg.Redis
	Rds = &redis.Pool{
		MaxIdle:     cfg.MaxIdleConn,
		MaxActive:   cfg.MaxOpenConn,
		IdleTimeout: time.Duration(cfg.ConnectTimeout) * time.Second,
		Wait:        false,
		Dial: func() (redis.Conn, error) {
			host := fmt.Sprintf("%v:%v", cfg.Network, cfg.Port)
			conn, err := redis.Dial(
				cfg.Network,
				host,
				redis.DialPassword(cfg.Pass),
				redis.DialDatabase(cfg.DbName),
				redis.DialConnectTimeout(time.Duration(cfg.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(cfg.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(cfg.WriteTimeout)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
