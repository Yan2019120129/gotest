package test_go_t

// 编译层示例：
// go:noinline 禁止编译器内联函数，常用于观察汇编、性能实验或防止测试被优化掉。
//
//go:noinline
func NoInlineAdd(a, b int) int {
	return a + b
}

// go:nosplit 禁止函数插入栈扩容检查，常见于 runtime 或极底层代码。
// 普通业务代码一般不需要使用；示例函数不调用其他函数，避免 nosplit 调用链过深。
//
//go:nosplit
func NoSplitAdd(a, b int) int {
	return a + b
}

// go:uintptrescapes 告诉编译器 uintptr 参数里可能保存了指针值。
// 这个指令常见于 syscall 或底层封装，用来避免指针过早被 GC 回收。
//
//go:uintptrescapes
func UintptrEscapesAddress(addr uintptr) uintptr {
	return addr
}
