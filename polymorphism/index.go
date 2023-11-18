package main

import "fmt"

// Animal 定义一个接口
type Animal interface {
	Speak() string
}

// Dog 定义Dog结构体
type Dog struct {
}

// Speak Dog结构体实现Animal接口的Speak方法
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat 定义Cat结构体
type Cat struct {
}

// Speak Cat结构体实现Animal接口的Speak方法
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// 创建一个Animal类型的切片，存放Dog和Cat的实例
	animals := []Animal{
		Dog{},
		Cat{},
	}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
