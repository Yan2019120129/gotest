//go:build !amd64

package test_go_t

// 非 amd64 平台保留同名函数，保证测试仍可编译。
func NoEscapeAdd(a, b int) int {
	return a + b
}
