package imp

import (
	"fmt"
)

type Mysql struct{}

func (dataBase *Mysql) Use() {
	fmt.Println("create mysql")
}

func (dataBase *Mysql) NewDatabase() {
	fmt.Println("create mysql")
}
