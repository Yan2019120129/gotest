package main

import (
	"fmt"
	"testing"
)

// TestGoroutine 测试协程的运作原理
func TestGoroutine(t *testing.T) {
	message := ""
	Instanct.Main()
	for {
		if _, err := fmt.Scanln(&message); err != nil {
			panic(err)
		}
		Instanct.Message <- message
	}
}
