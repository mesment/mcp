# 第一阶段：构建阶段
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum* ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server/main.go

# 第二阶段：运行阶段
FROM alpine:latest

# 安装基本工具，因为这是一个终端服务器，需要基本的shell命令
RUN apk --no-cache add ca-certificates tzdata bash

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/server .

# 确保README.md被正确复制到镜像中
COPY --from=builder /app/README.md .

# 为了调试目的，显示文件是否存在
RUN ls -la /root/

# 暴露端口（如果需要）
# EXPOSE 8080

# 设置时区为Asia/Shanghai
ENV TZ=Asia/Shanghai

# 运行应用
CMD ["./server"]