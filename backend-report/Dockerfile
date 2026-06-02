FROM golang:1.19-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod init sausage-backend-report 2>/dev/null || true && \
    go get github.com/gorilla/mux && \
    go get go.mongodb.org/mongo-driver/mongo && \
    go get go.mongodb.org/mongo-driver/mongo/options && \
    CGO_ENABLED=0 GOOS=linux go build -o backend-report .

FROM alpine:latest

COPY --from=0 /app/backend-report /backend-report

EXPOSE 8081

CMD ["/backend-report"]
