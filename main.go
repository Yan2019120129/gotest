package main

import "fmt"

const (
	serverAddr = "./product"
)

func main() {
	name := "/crawling/describe/images_I_41ljs+T-dO14S._SL1500_.jpg"
	tmpName := []byte(name)
	isPoint := false
	for i := len(tmpName) - 1; i > 0; i-- {
		if tmpName[i] == '+' || tmpName[i] == '-' {
			tmpName[i] = '_'
		}
		if tmpName[i] == '.' && isPoint {
			tmpName = append(tmpName[:i], tmpName[i+1:]...)
		}
		if tmpName[i] == '.' {
			isPoint = true
		}
	}
	fmt.Println(string(tmpName))
}
