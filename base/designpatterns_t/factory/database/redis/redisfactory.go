package redis

import (
	"gotest/base/designpatterns_t/factory/database/redis/imp"
)

func CreateDatabase() *imp.Redis {
	return &imp.Redis{}
}
