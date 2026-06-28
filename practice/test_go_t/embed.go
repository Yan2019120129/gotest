package test_go_t

import _ "embed"

// go:embed 在编译期把文件内容打进最终二进制。
//
//go:embed embed_assets/message.txt
var EmbeddedMessage string
