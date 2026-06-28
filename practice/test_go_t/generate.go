package test_go_t

// 源码层示例：
// go:generate 在编译前不会自动执行，需要手动运行 go generate。
// 这里模拟“业务配置生成代码”的场景，把 config.txt 中的应用名称生成到 zz_generated_config.go。
//
// 运行命令：
// go generate ./practice/test_go_t
//
//go:generate go run ./cmd/genconfig -in config.txt -out zz_generated_config.go -pkg test_go_t
