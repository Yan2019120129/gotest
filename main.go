package main

import (
	"fmt"
	"regexp"
)

func main() {
	v := "yanjiasjig-----123abc456"
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(v, -1)
	fmt.Println(numbers) // 输出: [123 456]
}
