package goto_t

import "fmt"

// GoTo 测试goto流程
func GoTo() {
	i := 0
run:
	fmt.Println("I:", i)
	i++
	if i == 30 {
		// continue 不可用
		// break 不可用
		return
	}
	goto run
}
