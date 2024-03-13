package int_t

import (
	"fmt"
)

type Test struct {
	OnMessage func(msg string) string
	gMessage  func(msg string)
}

var Instance = &Test{
	OnMessage: nil,
	gMessage:  nil,
}

func (t *Test) ForMessage(msg ...string) {
	for _, v := range msg {
		t.OnMessage(v)
	}
}

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
