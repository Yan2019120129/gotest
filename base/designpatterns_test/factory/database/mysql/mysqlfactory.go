package mysql

import (
	"gotest/designpatterns/factory/database"
	"gotest/designpatterns/factory/database/mysql/imp"
)

// MysqlFactory 具体工厂类型
type MysqlFactory struct{}

func (f *MysqlFactory) CreateDatabase() database.Database {
	return &imp.Mysql{}
}
