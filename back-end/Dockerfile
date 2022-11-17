FROM golang:1.19-alpine as builder
WORKDIR /usr/local/bin
# 使用国内源
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
# 预处理依赖
COPY go.mod go.sum ./
RUN go mod download && go mod verify
# 编译二进制文件
COPY . .
RUN go build -v -o /usr/local/bin/app .

# 使用轻量级镜像运行二进制文件
FROM alpine:latest
COPY config/* ./config/
COPY --from=builder /usr/local/bin/app .
EXPOSE 8000
ENTRYPOINT [ "./app", "-mode=release"]