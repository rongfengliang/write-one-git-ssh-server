FROM golang:1.17-alpine AS build-env
WORKDIR /go/src/app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

ENV  GO111MODULE=on
ENV  GOPROXY=https://goproxy.cn
COPY . .
RUN apk update && apk add git \
    && go build -o git-http

FROM alpine:latest
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add ca-certificates git-daemon && rm -rf /var/cache/apk/*
COPY --from=build-env /go/src/app/git-http /app/git-http
ENTRYPOINT [ "/app/git-http" ]