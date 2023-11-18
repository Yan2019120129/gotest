package database

import "fmt"

type Oracle struct{}

func (database *Oracle) Connect() {
	fmt.Println("create oracle")
}
