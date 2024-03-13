package mysql

import (
	"gotest/base/designpatterns_t/factory/database/mysql/imp"
)

// MysqlFactory 具体工厂类型
type MysqlFactory struct{}

func (f *MysqlFactory) CreateDatabase() *imp.Mysql {
	return &imp.Mysql{}
}
