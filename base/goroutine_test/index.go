package main

import (
	"fmt"
)

var Instanct = &Test{
	Message: make(chan string, 50),
}

// Test
type Test struct {
	Message chan string
}

// test1 测试select获取通道数据方式
func (i *Test) test1() {
	fmt.Println("test1开启")
	defer fmt.Println("test1关闭")
	msg := ""
	select {
	case msg = <-i.Message:
		fmt.Println("test1 message:", msg)
		if msg == "break" {
			fmt.Println("test1 break")
			break
		}
		if msg == "return" {
			fmt.Println("test1 return")
			return
		}
		fmt.Println("test1 msg:", msg)
	}
}

// test2 测试select获取通道数据方式
func (i *Test) test2() {
	fmt.Println("test2开启")
	defer fmt.Println("test2关闭")
	msg := ""
	select {
	case msg = <-i.Message:
		fmt.Println("test2 message:", msg)
		if msg == "break" {
			fmt.Println("test2 break")
			break
		}
		if msg == "return" {
			fmt.Println("test2 return")
			return
		}
		fmt.Println("test2 msg:", msg)
	}
}

// test3 测试select获取通道数据方式
func (i *Test) test3() {
	fmt.Println("test3开启")
	defer fmt.Println("test3关闭")
	msg := ""
	select {
	case msg = <-i.Message:
		fmt.Println("test3 message:", msg)
		if msg == "break" {
			fmt.Println("test3 break")
			break
		}
		if msg == "return" {
			fmt.Println("test3 return")
			return
		}
		fmt.Println("test3 msg:", msg)
	default:
		fmt.Printf("no communication\n")
	}
}

// test4 测试 for 获取通道数据方式
func (i *Test) test4() {
	for {
		msg := <-i.Message
		fmt.Println("test4 msg:", msg)
	}
}

// test5 测试 for 获取通道数据方式
func (i *Test) test5() {
	for {
		msg := <-i.Message
		fmt.Println("test5 msg:", msg)
	}
}

// test6 测试 for 获取通道数据方式
func (i *Test) test6() {
	for {
		msg := <-i.Message
		fmt.Println("test6 msg:", msg)
	}
}

// close 关闭线程
func (i *Test) close() {

}

// Main 主方法  default 会持续遍历, select 各个协程之间不会同时收到消息
func (i *Test) Main() {
	go i.test1()
	go i.test2()
	go i.test3()
	for {
		message := ""
		if _, err := fmt.Scanln(&message); err != nil {
			panic(err)
		}

		Instanct.Message <- message
	}
}

// Main1 测试各个协程会不会同时收到消息
func (i *Test) Main1() {
	go i.test4()
	go i.test5()
	go i.test6()
	for {
		message := ""
		if _, err := fmt.Scanln(&message); err != nil {
			panic(err)
		}

		Instanct.Message <- message
	}
}

func main() {
	//Instanct.Main()
	Instanct.Main1()
}
