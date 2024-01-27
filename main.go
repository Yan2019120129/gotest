package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("systemctl:", runtime.GOOS)
}
