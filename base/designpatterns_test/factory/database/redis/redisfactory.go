package redis

import (
	"gotest/base/designpatterns_test/factory/database/redis/imp"
)

func CreateDatabase() *imp.Redis {
	return &imp.Redis{}
}
