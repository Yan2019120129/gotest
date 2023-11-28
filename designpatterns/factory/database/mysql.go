package database

import (
	"fmt"
)

type Mysql struct{}

func (dataBase *Mysql) Connect() {
	fmt.Println("create mysql")
}
