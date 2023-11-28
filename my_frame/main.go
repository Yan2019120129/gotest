package main

import (
	"errors"
	models "gotest/my_frame/ models"
	"gotest/my_frame/database"
)

func main() {
	factory := new(database.New)
	db := factory.NewMysql().Connect()
	user := new(models.User)
	result := db.Find(&user)
	if result.RowsAffected == 0 {
		errors.New("")
	}
}
