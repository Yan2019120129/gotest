package main

import "fmt"

func main() {
	message := "yanjiajie"
	for i, v := range message {
		fmt.Println("i:", i)
		fmt.Println("v:", string(v))
	}
}
