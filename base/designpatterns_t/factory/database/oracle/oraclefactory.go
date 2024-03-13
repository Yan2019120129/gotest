package oracle

import (
	"gotest/base/designpatterns_t/factory/database/oracle/imp"
)

// OracleFactory 具体工厂类型
type OracleFactory struct{}

func (of *OracleFactory) CreateDatabase() *imp.Oracle {
	return &imp.Oracle{}
}
