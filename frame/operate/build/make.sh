#!/bin/bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o op ./main.go

cp op ./build/op

# 上传文件并捕获返回的 JSON
response=$(curl -s http://tw10b0135.onething.net:9999/v1/files --form upload=@./build/op)

# 提取 URL
url=$(echo "$response" | grep -oP '(http://[^"]+)')

# 输出 URL
echo "$url"