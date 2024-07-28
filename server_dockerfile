# Server Dockerfile
FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN go build -o server server.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

CMD ["./server"]
