package secret

// HiddenToken 是未导出的内部函数，模拟旧模块中暂时不想公开的能力。
func HiddenToken(userID string) string {
	return "token:" + userID
}
