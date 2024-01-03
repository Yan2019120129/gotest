package test_init

import "fmt"

// Test  测试用的结构体
type Test struct {
	name string // 名字
	age  int    // 年龄
}

func (t *Test) Name() string {
	return t.name
}

func (t *Test) SetName(name string) {
	fmt.Println("name:", name)
	t.name = name
}

func (t *Test) Age() int {
	return t.age
}

func (t Test) SetAge(age int) {
	fmt.Println("age:", age)
	t.age = age
}
