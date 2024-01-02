package main

import (
	"fmt"
)

func main() {
	IntSum()
}

// IntSum 查找中间值
func IntSum() {
	//a := (2 << 62) - 300
	//c := (2 << 62) - 100

	a := 301
	c := 101

	b := (a + c) / 2
	//b := a/2 + c/2

	d := a + (c-a)/2

	//d := a + c/2 - a/2
	fmt.Printf("a:%v\n", a)
	fmt.Printf("c:%v\n", c)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("d:%v\n", d)

	//a := 13982414433
	//c := 17982414433
	//b := (a + c) / 2
	//fmt.Println("data:", b)
}
