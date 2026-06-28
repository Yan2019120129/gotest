//go:build amd64

package test_go_t

// go:noescape 只能用于没有 Go 函数体、由汇编实现的函数声明。
// 它告诉编译器：传入该函数的指针参数不会逃逸到堆上。
// 这里用简单加法演示声明和汇编实现的连接方式。
//
//go:noescape
func NoEscapeAdd(a, b int) int
