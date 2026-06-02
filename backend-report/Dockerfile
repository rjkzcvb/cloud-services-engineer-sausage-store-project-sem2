FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod init backend-report 2>/dev/null || true

COPY *.go ./

RUN go build -o backend-report .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /app/backend-report .

EXPOSE 8081

CMD ["./backend-report"]
