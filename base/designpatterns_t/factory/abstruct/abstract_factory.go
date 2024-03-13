package abstruct

import (
	"gotest/base/designpatterns_t/factory/log"
	"gotest/common/module/gorm/database"
)

type AbstractFactory interface {
	CreateDatabase() database.Database
	CreateLog() log.Log
}
