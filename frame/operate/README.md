# gotest
go语言基础练习例子

```shell # 编译程序
GOOS=linux GOARCH=amd64 go build -o main
```

```shell
# ldflags 去除调试信息  -s：省略符号表和调试信息    -w：禁用 DWARF 生成（进一步减小体积）
CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o op main.go

[[ -d ./build ]] || mkdir ./build
[[ -f ./build/op ]] && rm -rf ./build/op
mv op ./build
```

```shell
timestamp=$(date "+%Y-%m-%d_%H:%M:%S")
mv /usr/bin/op /tmp/op_${timestamp}
cp  ./build/op /usr/bin/
```

```shell
[[ -f /usr/bin/op ]] && sudo rm -f /usr/bin/op
sudo mv ./build/op /usr/bin/
```