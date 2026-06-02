FROM golang:1.19-alpine AS builder

WORKDIR /app

# Устанавливаем git (нужен для go mod)
RUN apk add --no-cache git

# Копируем все файлы
COPY . .

# Инициализируем go module
RUN go mod init sausage-backend-report 2>/dev/null || true

# Устанавливаем зависимости
RUN go get github.com/gorilla/mux 2>/dev/null || true
RUN go get go.mongodb.org/mongo-driver/mongo 2>/dev/null || true
RUN go get go.mongodb.org/mongo-driver/mongo/options 2>/dev/null || true

# Скачиваем зависимости
RUN go mod tidy

# Показываем файлы для отладки
RUN ls -la && echo "=== Go files ===" && ls -la *.go

# Собираем приложение (с выводом ошибок)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o backend-report .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend-report .

EXPOSE 8081

CMD ["./backend-report"]
