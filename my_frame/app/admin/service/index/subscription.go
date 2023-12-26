package index

import "gotest/my_frame/module/redis"

// SubRds 订阅信息
func SubRds() (interface{}, error) {
	rdsConn := redis.RdsPubSubConn
	defer rdsConn.Close()
	rdsConn.Subscribe()
	return nil, nil
}
