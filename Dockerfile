# 构建阶段
FROM golang:1.18-alpine AS builder

WORKDIR /build

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

# 打包阶段
FROM alpine:latest

LABEL maintainer="simon"

# 设置永久环境变量, 或 ENV key=value
ENV WORKDIR /data/app

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 添加应用可执行文件，并设置执行权限
COPY --from=builder /build/main  $WORKDIR/main
RUN chmod +x $WORKDIR/main

EXPOSE 9999/tcp

WORKDIR $WORKDIR

CMD ["./main"]