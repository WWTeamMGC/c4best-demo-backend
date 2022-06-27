FROM golang:1.17 as builder

##
## Build
##
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o c4best_demo ./cmd/

##
## Build
##
FROM alpine:3.14

WORKDIR /root/
RUN apk add --no-cache ffmpeg
COPY --from=builder /app/c4best_demo ./c4best_demo

EXPOSE 8080

ENTRYPOINT ["./c4best_demo"]