package int_test

import (
	"fmt"
)

// IntSum 查找中间值
func IntSum() {
	a := 301
	c := 101
	b := a + (c-a)/2

	fmt.Printf("a:%v\n", a)
	fmt.Printf("c:%v\n", c)
	fmt.Printf("b:%v\n", b)
}

// IntSumOverflow 查找中间值溢出的情况
func IntSumOverflow() {
	a := (2 << 62) - 300
	c := (2 << 62) - 100

	b := (a + c) / 2

	fmt.Printf("a:%v\n", a)
	fmt.Printf("c:%v\n", c)
	fmt.Printf("b:%v\n", b)
}
