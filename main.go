package main

import "fmt"

func main() {
	message := ""
	fmt.Print("等待输入：")
	_, err := fmt.Scan(&message)
	if err != nil {
		fmt.Println("发生错误:", err)
	}
	fmt.Println("输入内容:", message)
}
