//go:build directive_demo

package test_go_t

// BuildTagMode 演示 go:build：运行 go test -tags directive_demo 时编译该文件。
func BuildTagMode() string {
	return "directive_demo"
}
