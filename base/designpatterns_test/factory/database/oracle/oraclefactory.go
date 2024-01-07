package oracle

import (
	"gotest/designpatterns/factory/database"
	"gotest/designpatterns/factory/database/oracle/imp"
)

// OracleFactory 具体工厂类型
type OracleFactory struct{}

func (of *OracleFactory) CreateDatabase() database.Database {
	return &imp.Oracle{}
}
