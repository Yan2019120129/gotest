package init_test

import (
	"fmt"
	"gotest/base/init_test/first_naming"
)

var Name = ""

func init() {
	Name = first_naming.FirstName + "jiajie"
	fmt.Println("我给了他起名字")
}

func Naming() {
	fmt.Println("我起好名字了：", Name)
}
