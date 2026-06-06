#!/bin/bash

set -e

# ========================
# 默认参数（可覆盖）
# ========================

# 编译入口文件
SRC_FILE=${1:-main.go}

# 输出二进制名称（可传第2个参数）
BINARY_NAME=${2:-main}

# 远程用户（可传第3个参数）
REMOTE_USER=${3:-root}

# 远程IP（可传第4个参数）
REMOTE_HOST=${4:-8.138.57.34}

# 远程目录（可传第5个参数）
REMOTE_PATH=${5:-/root/file/shell}

# ========================
# 时间戳（精确到分钟）
# ========================
TIME=$(date +"%Y%m%d_%H%M")

FINAL_BINARY="${BINARY_NAME}_${TIME}"

# ========================
# 检查源文件
# ========================
if [ ! -f "$SRC_FILE" ]; then
  echo "❌ source file not found: $SRC_FILE"
  exit 1
fi

echo "=============================="
echo "📄 source file : $SRC_FILE"
echo "📦 binary name  : $FINAL_BINARY"
echo "👤 remote user  : $REMOTE_USER"
echo "🌐 remote host  : $REMOTE_HOST"
echo "📁 remote path  : $REMOTE_PATH"
echo "=============================="
echo "./build.sh $SRC_FILE $FINAL_BINARY $REMOTE_USER $REMOTE_HOST $REMOTE_PATH"

# ========================
# 编译
# ========================
echo "🔧 building..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -ldflags="-s -w" -o ${FINAL_BINARY} ${SRC_FILE}

echo "✅ build success: ${FINAL_BINARY}"

# ========================
# 上传
# ========================
echo "🚀 uploading..."

scp ${FINAL_BINARY} ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_PATH}

echo "🎉 done"