package main

import (
	"fmt"
	"gotest/crawling/test"
)

func main() {
	ts := test.NewTestStruct()
	fmt.Println("ts:", ts)
	ts.GetName("yan")
	t := new(test.TestStruct)
	fmt.Println("t:", t)
}
