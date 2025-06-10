package main

import (
	"business/utils"
	"fmt"
)

func main() {
	val := utils.GetBZInstanceCount()
	fmt.Println(val)
}
