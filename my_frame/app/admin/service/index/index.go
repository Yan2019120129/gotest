package index

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gotest/my_frame/module/cache"
)

func Index() (interface{}, error) {
	rds := cache.RdsPool.Get()

	defer rds.Close()
	_, err := rds.Do("Set", "age", 18)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := redis.Int(rds.Do("GET", "age"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return nil, err
	}

	return data, nil
}
