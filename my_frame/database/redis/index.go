package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var Rds *redis.Pool

// InitRedis 初始化Redis
func InitRedis() {
	Rds = &redis.Pool{
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: time.Duration(30) * time.Second,
		Wait:        false,
		Dial: func() (redis.Conn, error) {
			host := "127.0.0.1:6379"
			conn, err := redis.Dial(
				"tcp",
				host,
				redis.DialPassword(""),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(time.Duration(30)*time.Second),
				redis.DialReadTimeout(time.Duration(30)*time.Second),
				redis.DialWriteTimeout(time.Duration(30)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
