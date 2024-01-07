package abstruct

import (
	"gotest/designpatterns/factory/database"
	"gotest/designpatterns/factory/log"
)

type AbstractFactory interface {
	CreateDatabase() database.Database
	CreateLog() log.Log
}
