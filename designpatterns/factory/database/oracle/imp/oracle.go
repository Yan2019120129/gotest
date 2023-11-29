package imp

import "fmt"

type Oracle struct{}

func (database *Oracle) Use() {
	fmt.Println("create oracle")
}

func (database *Oracle) NewDatabase() {
	fmt.Println("create oracle")
}
