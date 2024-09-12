FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
