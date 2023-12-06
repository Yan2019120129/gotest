package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/my_frame/models"
	"time"
)

var Rds *redis.Pool

// Init 初始化Redis
func Init(cfg *models.RedisConfig) {
	Rds = &redis.Pool{
		MaxIdle:     cfg.MaxIdleConn,
		MaxActive:   cfg.MaxOpenConn,
		IdleTimeout: time.Duration(cfg.ConnectTimeout) * time.Second,
		Wait:        false,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				cfg.Network,
				getDsn(cfg),
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

func getDsn(cfg *models.RedisConfig) string {
	return fmt.Sprintf("%v:%v", cfg.Network, cfg.Port)
}
