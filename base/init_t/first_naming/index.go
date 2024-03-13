package first_naming

import "fmt"

var FirstName = ""

func init() {
	FirstName = "yan"
	fmt.Println("我开始给她起姓氏了")
}

func Naming() {
	fmt.Println("我知道他姓什么了：", FirstName)
}
