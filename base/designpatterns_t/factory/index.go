package main

import (
	"gotest/base/designpatterns_t/factory/database/mysql"
	"gotest/base/designpatterns_t/factory/database/oracle"
)

func main() {
	_mysql := new(mysql.MysqlFactory).CreateDatabase()
	_mysql.Use()

	_oracle := new(oracle.OracleFactory).CreateDatabase()
	_oracle.Use()
}
