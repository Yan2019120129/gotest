package main

import "fmt"

func main() {
	a := make([]int, 5)

	a[0] = 1
	a[1] = 2
	//a[2] = 3
	//a[3] = 4
	//a = append(a, 1)
	//a = append(a, 2)
	//a = append(a, 3)
	//a = append(a, 4)
	fmt.Println("a:", len(a))
}
