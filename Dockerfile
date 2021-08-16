FROM golang:1.16-alpine AS builder
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"
COPY . /build
WORKDIR /build
#RUN go mod tidy
#RUN go get -u github.com/swaggo/swag/cmd/swag
#RUN swag init --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
RUN sed -i "s@https://dl-cdn.alpinelinux.org/@https://mirrors.huaweicloud.com/@g" /etc/apk/repositories
RUN apk update && apk add upx
RUN upx codeql-ct

FROM alpine
ENV WAIT_VERSION 2.7.3
# ENV WAIT_RELEASE https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait
ENV WAIT_RELEASE https://st0n3-dev.obs.cn-south-1.myhuaweicloud.com/docker-compose-wait/release/$WAIT_VERSION/wait
ADD $WAIT_RELEASE /wait
RUN chmod +x /wait

RUN mkdir -p /app
COPY --from=builder /build/codeql-ct /app/

WORKDIR /app
CMD /wait && ./codeql-ct