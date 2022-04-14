FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# 安装ca-certificates，发送HTTPS请求，否则会报错"x509: certificate signed by unknown authority"
RUN apk update && apk upgrade
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache build-base=0.5-r2
RUN update-ca-certificates

# 编译项目
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o server .

# 二阶段编译
FROM alpine AS releaser

WORKDIR ./
COPY --from=builder /build/server ./server
COPY --from=builder /build/resource ./resource
COPY --from=builder /build/config.docker.yaml ./config.docker.yaml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/shanghai" > /etc/timezone;

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
