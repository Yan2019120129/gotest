# gotest
go语言基础练习例子

```shell # 编译程序
GOOS=linux GOARCH=amd64 go build -o main
```

```shell
# ldflags 去除调试信息    -s：省略符号表和调试信息    -w：禁用 DWARF 生成（进一步减小体积）
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o networkStats
```