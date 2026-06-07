#!/bin/bash

set -e

# 时间戳（精确到分钟）
TIME=$(date +"%Y%m%d_%H%M")

# 文件名
BINARY="priceMonitor_${TIME}"

echo "🔧 building..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINARY} main.go

echo "📦 built: ${BINARY}"

echo "🚀 uploading to server..."

scp ${BINARY} root@8.138.57.34:/root/file/priceMonitor/history/

echo "✅ done"