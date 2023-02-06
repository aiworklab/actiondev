FROM golang:1.18 AS builder
LABEL maintainer="aichy"

# 设置时区环境变量
ENV TZ Asia/Shanghai
ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV GOPROXY https://goproxy.cn,direct
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}

# Stage 1 Build
COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
ENV CGO_ENABLED=0
RUN pwd
RUN go build -o main
RUN mkdir /data
RUN cp main /data/
RUN cp config.toml /data/

FROM alpine:latest
ENV TZ "Asia/Shanghai"

COPY --from=builder /data/ /data/
EXPOSE 8100
ENTRYPOINT ["/data/main","-c","/data/config.toml"]
