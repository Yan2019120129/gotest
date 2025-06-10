#!/bin/bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o custom_switch /home/yan/Documents/file/gofile/gotest/middleware/bandwidth_summary/main.go

mkdir -p ./build/custom_switch/conf

cp custom_switch ./build/custom_switch

cp ./conf/config.yaml ./build/custom_switch/conf

cd ./build
tar -czvf custom_switch.tar.gz ./custom_switch

rm -rf ./build/custom_switch

# 上传文件并捕获返回的 JSON
response=$(curl -s http://tw10b0135.onething.net:9999/v1/files --form upload=@./custom_switch.tar.gz)

# 提取 URL
url=$(echo "$response" | grep -oP '(http://[^"]+)')

# 输出 URL
echo "$url"