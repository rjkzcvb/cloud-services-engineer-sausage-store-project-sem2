FROM golang:1.19-alpine AS builder

WORKDIR /app

# Устанавливаем git (нужен для go mod)
RUN apk add --no-cache git

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходники
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o backend-report .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend-report .

EXPOSE 8081

CMD ["./backend-report"]
