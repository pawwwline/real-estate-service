FROM golang:1.22 AS builder

WORKDIR /real-estate-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o real-estate-service ./cmd/real-estate-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /real-estate-service/real-estate-service .

COPY config/local.yaml /root/config.yaml

COPY internal/db/migrations /root/migrations

EXPOSE 8080

ENV CONFIG_PATH=/root/config.yaml

CMD ["./real-estate-service"]
