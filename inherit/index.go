package inherit

import "fmt"

type Inherit struct {
	name string
	age  int64
}

func NewInherit(name string, age int64) *Inherit {
	return &Inherit{name: name, age: age}
}

func (i *Inherit) Name() *Inherit {
	return i
}

func (i *Inherit) SetName(name string) {
	i.name = name
}

func (i *Inherit) Age() *Inherit {
	return i
}

func (i *Inherit) SetAge(age int64) {
	i.age = age
}

func (i *Inherit) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d", i.name, i.age)
}
