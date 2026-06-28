package test_go_t

import (
	_ "gotest/practice/test_go_t/internal/secret"
	_ "unsafe"
)

// LinknameHiddenToken 演示 go:linkname 访问其他包的未导出函数。
// 该指令会绕过 Go 的包可见性规则，真实项目应优先使用正常导出的 API。
func LinknameHiddenToken(userID string) string {
	return hiddenToken(userID)
}

//go:linkname hiddenToken gotest/practice/test_go_t/internal/secret.HiddenToken
func hiddenToken(userID string) string
