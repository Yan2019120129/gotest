package redis

import (
	"gotest/designpatterns/factory/database"
	"gotest/designpatterns/factory/database/redis/imp"
)

func CreateDatabase() database.Database {
	return &imp.Redis{}
}
