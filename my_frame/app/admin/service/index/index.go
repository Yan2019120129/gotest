package index

import (
	"fmt"
	"gotest/my_frame/module/cache"
)

func Index() (interface{}, error) {
	rds := cache.RdsPool.Get()

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
