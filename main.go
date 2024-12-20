package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var val any = "2programmer"
	tmpVal, _ := json.Marshal(val)
	fmt.Println(string(tmpVal))
	var val1 = "2programmer"
	fmt.Println(val1)
}
