//go:build ignore

package examples_disabled

// go:inline 不是常规业务代码中稳定可依赖的优化开关。
// Go 编译器通常由自身启发式规则决定是否内联；业务代码一般只使用 go:noinline 禁止内联。
// 这里保留语法示意，使用 ignore 构建标签避免影响正常测试。
//
//go:inline
func inlineCandidate(a, b int) int {
	return a + b
}
