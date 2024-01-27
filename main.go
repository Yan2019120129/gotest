package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Struct:", Test)
	fmt.Println("Struct:", Message{})

	Test.Age = 18
	Test.Name = "yan"
	marshal, err := json.Marshal(Test)
	if err != nil {
		return
	}

	fmt.Println("data:", string(marshal))
}

var Test struct {
	Name string
	Age  int
}

type Message struct {
	Email int
	Name  string
}
