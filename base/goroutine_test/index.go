package main

import (
	"context"
	"fmt"
)

// Test
type Test struct {
	Message chan string
	ctx     context.Context
	Close   context.CancelFunc
	fun     func(parent context.Context) (ctx context.Context, cancel context.CancelFunc)
}

func NewInstance() *Test {
	ctx, off := context.WithCancel(context.Background())
	return &Test{
		Message: make(chan string, 50),
		ctx:     ctx,
		Close:   off,
		fun:     context.WithCancel,
	}
}

// test1 测试select获取通道数据方式
func (i *Test) test1() {
	fmt.Println("test1开启")
	defer fmt.Println("test1关闭")
	for {
		select {
		case msg, isClose := <-i.Message:
			if !isClose {
				return
			}
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
}

// test2 测试select获取通道数据方式
func (i *Test) test2() {
	fmt.Println("test2开启")
	defer fmt.Println("test2关闭")
	for {
		select {
		case msg, isClose := <-i.Message:
			if !isClose {
				return
			}
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
}

// test3 测试select获取通道数据方式
func (i *Test) test3() {
	fmt.Println("test3开启")
	defer fmt.Println("test3关闭")
	for {
		select {
		case msg, isClose := <-i.Message:
			if !isClose {
				return
			}
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

		}
	}
}

// test4 测试 for 获取通道数据方式
func (i *Test) test4() {
	defer fmt.Println("test4 close")
	for {
		select {
		case <-i.ctx.Done():
			return
		case msg := <-i.Message:
			fmt.Println("test4 msg:", msg)
		}
	}
}

// test5 测试 for 获取通道数据方式
func (i *Test) test5() {
	defer fmt.Println("test5 close")
	for {
		select {
		case <-i.ctx.Done():
			return
		case msg := <-i.Message:
			fmt.Println("test5 msg:", msg)
		}
	}
}

// test6 测试 for 获取通道数据方式
func (i *Test) test6() {
	defer fmt.Println("test6 close")
	for {
		select {
		case <-i.ctx.Done():
			return
		case msg := <-i.Message:
			fmt.Println("test6 msg:", msg)
		}
	}
}

// Run 关闭线程
func (i *Test) Run() *Test {
	//go i.test1()
	//go i.test2()
	//go i.test3()
	go i.test4()
	go i.test5()
	go i.test6()
	return i
}

// Main 主方法  default 会持续遍历, select 各个协程之间不会同时收到消息
func Main() {
	i := NewInstance()
	go i.test1()
	go i.test2()
	go i.test3()
	for {
		message := ""
		if _, err := fmt.Scanln(&message); err != nil {
			panic(err)
		}
		if message == "close" {
			close(i.Message)
			return
		}

		i.Message <- message
	}
}

// Main1 测试各个协程会不会同时收到消息
func Main1() {
	i := NewInstance().Run()
	for {
		message := ""
		if _, err := fmt.Scanln(&message); err != nil {
			panic(err)
		}

		if message == "close" {
			i.Close()
		}

		if message == "run" {
			i = NewInstance().Run()
		}

		i.Message <- message
	}
}

func main() {
	//Main()
	Main1()
}
