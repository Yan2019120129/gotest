package main

import (
	"gotest/designpatterns/factory/database/mysql"
	"gotest/designpatterns/factory/database/oracle"
)

func main() {
	_mysql := new(mysql.MysqlFactory).CreateDatabase()
	_mysql.Use()

	_oracle := new(oracle.OracleFactory).CreateDatabase()
	_oracle.Use()
}
