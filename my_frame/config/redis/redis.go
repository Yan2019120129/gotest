package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/my_frame/config"
	"gotest/my_frame/models"
	"sync"
	"time"
)

// 初始化redis连接池,保证全局单例
var once sync.Once

var Rds *redis.Pool

//func init() {
//	Init()
//}

// Init 初始化Redis
func Init() {
	if Rds == nil {
		once.Do(func() {
			cfg := config.GetRedis()
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
						panic(err)
						return nil, err
					}
					return conn, nil
				},
			}
			fmt.Printf("内存地址：%p----->Res实例创建成功！！！\n", Rds)
		})
	} else {
		fmt.Println("已经存在Res实例！！！")
	}
}

func getDsn(cfg *models.RedisConfig) string {
	return fmt.Sprintf("%v:%v", cfg.Network, cfg.Port)
}
