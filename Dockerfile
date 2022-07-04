FROM golang:1.17 as builder
MAINTAINER c4best

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
##
## Build
##
WORKDIR /app

COPY go.mod go.sum ./

COPY . .
COPY --from=build /app/config.yaml /config.yaml
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

##
## Build
##
FROM alpine:3.14

WORKDIR /root/
COPY --from=builder /app/c4best_demo ./c4best_demo

EXPOSE 8080

ENTRYPOINT ["./c4best_demo"]