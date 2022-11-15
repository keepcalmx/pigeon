FROM golang:1.19-alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /usr/local/bin

# 预处理依赖
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

ENTRYPOINT [ "/usr/local/bin/app", "-mode=release"]