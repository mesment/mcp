#!/bin/bash

set -e

# 定义镜像名称和标签
IMAGE_NAME="mcp-terminal-server"
IMAGE_TAG="latest"

echo "构建Docker镜像: ${IMAGE_NAME}:${IMAGE_TAG}"
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

echo "\n镜像构建完成!"
echo "运行镜像的命令示例:"
echo "docker run -i --rm ${IMAGE_NAME}:${IMAGE_TAG}"