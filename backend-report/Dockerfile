FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum (если есть)
COPY go.* ./
RUN if [ -f go.mod ]; then go mod download; fi

# Копируем исходники
COPY *.go ./

# Создаем go.mod если нет
RUN if [ ! -f go.mod ]; then \
        go mod init sausage-backend-report && \
        go get github.com/gorilla/mux && \
        go get go.mongodb.org/mongo-driver/mongo && \
        go get go.mongodb.org/mongo-driver/mongo/options && \
        go mod tidy; \
    fi

# Собираем
RUN CGO_ENABLED=0 GOOS=linux go build -o backend-report .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend-report .

EXPOSE 8081

CMD ["./backend-report"]
