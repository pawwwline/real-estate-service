FROM golang:1.22 AS builder

WORKDIR /real-estate-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o real-estate-service ./cmd/real-estate-service

EXPOSE 8080

ENV CONFIG_PATH=/real-estate-service/config/docker.yaml

CMD ["./real-estate-service"]
