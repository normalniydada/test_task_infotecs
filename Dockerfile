# Используем официальный образ Go для сборки
FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/internal/config/config.yaml ./internal/config/config.yaml

RUN chmod +x main

# Команда для запуска
CMD ["./main"]