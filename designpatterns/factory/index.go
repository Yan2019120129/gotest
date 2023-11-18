package main

import "gotest/designpatterns/factory/database"

func main() {
	_Mysql := new(database.MysqlFactory)
	mysql := _Mysql.CreateDatabase()
	mysql.Connect()
}
