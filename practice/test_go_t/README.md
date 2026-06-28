# Go Directive 示例

本目录按“源码层、编译层、运行时层”整理 Go directive 和 runtime 行为示例。

## 运行方式

生成代码：

```bash
go generate ./practice/test_go_t
```

运行默认测试：

```bash
go test ./practice/test_go_t/...
```

运行带 `go:build` 标签的测试：

```bash
go test -tags directive_demo ./practice/test_go_t/...
```

## 已覆盖内容

- `//go:generate`：见 `generate.go`，读取 `config.txt` 生成 `zz_generated_config.go`。
- `//go:build`：见 `buildtag_default.go` 和 `buildtag_demo.go`，通过 `-tags directive_demo` 切换编译文件。
- `//go:embed`：见 `embed.go`，把 `embed_assets/message.txt` 编译进变量。
- `//go:noinline`：见 `compiler.go`，禁止函数内联。
- `//go:nosplit`：见 `compiler.go`，禁止插入栈扩容检查。
- `//go:noescape`：见 `noescape_amd64.go` 和 `noescape_amd64.s`，声明由汇编实现的函数。
- `//go:uintptrescapes`：见 `compiler.go`，说明 uintptr 参数中可能保存指针。
- `//go:linkname`：见 `linkname.go`，访问 `internal/secret` 中未导出的函数。

## 仅说明，不参与编译

`examples_disabled` 下的文件使用 `//go:build ignore`，用于保留语法和中文说明，不参与正常编译：

- `//go:inline`：Go 编译器通常由启发式规则控制内联，业务代码不要依赖它。
- `//go:nowritebarrier`
- `//go:nowritebarrierrec`
- `//go:yeswritebarrierrec`

写屏障相关 directive 属于 Go runtime 内部约束，普通包直接使用会被编译器拒绝。
