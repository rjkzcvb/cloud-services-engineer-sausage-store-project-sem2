FROM golang:1.19-alpine AS builder

WORKDIR /app

# Копируем все .go файлы
COPY *.go ./

# Инициализируем go module если нет
RUN if [ ! -f go.mod ]; then \
        go mod init sausage-backend-report; \
        go get github.com/gorilla/mux; \
        go get go.mongodb.org/mongo-driver/mongo; \
        go get go.mongodb.org/mongo-driver/mongo/options; \
        go mod tidy; \
    fi

# Скачиваем зависимости
RUN go mod download

# Собираем
RUN CGO_ENABLED=0 GOOS=linux go build -o backend-report .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend-report .

EXPOSE 8081

CMD ["./backend-report"]
