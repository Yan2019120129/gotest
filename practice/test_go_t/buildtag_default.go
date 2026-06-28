//go:build !directive_demo

package test_go_t

// BuildTagMode 演示 go:build：默认没有传入 directive_demo 标签时编译该文件。
func BuildTagMode() string {
	return "default"
}
