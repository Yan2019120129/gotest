package index

import (
	"fmt"
	"gotest/my_frame/config/redis"
)

func Index() (interface{}, error) {
	rds := redis.Rds.Get()

	defer rds.Close()
	_, err := rds.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := rds.Do("Get", "abc")
	if err != nil {
		fmt.Println("get abc failed,", err)
		return nil, err
	}

	return data, nil
}
