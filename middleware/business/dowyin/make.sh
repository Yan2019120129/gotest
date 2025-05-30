#!/bin/bash
# 构建二进制文件
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o task_dy_limit_bw

# 上传文件并捕获返回的 JSON
response=$(curl -s http://tw10b0135.onething.net:9999/v1/files --form upload=@task_dy_limit_bw)

# 提取 URL
url=$(echo "$response" | grep -oP '(http://[^"]+)')

# 输出 URL
echo "$url"