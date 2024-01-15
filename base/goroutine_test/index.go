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

// test1
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
	}
}

// test2
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
	}
}

// test3
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
	default:
		fmt.Printf("no communication\n")
	}
}

// close 关闭线程
func (i *Test) close() {

}

// Main 主方法
func (i *Test) Main() {
	go i.test1()
	go i.test2()
	go i.test3()
}

func main() {
	message := ""
	Instanct.Main()
	for {
		if _, err := fmt.Scanln(&message); err != nil {
			panic(err)
		}
		if message == "start" {
			Instanct.Main()
			continue
		}
		Instanct.Message <- message
	}
}
