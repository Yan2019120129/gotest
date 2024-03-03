package indexserver

import (
	"github.com/gomodule/redigo/redis"
	"my-frame/app/admin/service/dto"
	"my-frame/module/cache"
)

// Rds 设置参数
func Rds(params *dto.RdsParams) (interface{}, error) {
	rdsconn := cache.RdsPool.Get()
	defer rdsconn.Close()

	switch params.Type {
	case 1:
		data, err := rdsconn.Do(params.Command, params.Key, params.Value)
		if err != nil {
			return nil, err
		}
		return data, nil
	case 2:
		data, err := redis.String(rdsconn.Do(params.Command, params.Key))
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, nil
	}
}
