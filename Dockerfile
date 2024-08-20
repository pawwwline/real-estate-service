FROM golang:1.22 AS builder

WORKDIR /real-estate-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o real-estate-service ./cmd/real-estate-service

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin


EXPOSE 8080

ENV CONFIG_PATH=/real-estate-service/config/docker.yaml

CMD ["sh", "-c", "migrate -path internal/db/migrations -database postgres://myuser:pass@db:5432/estatedb?sslmode=disable up && ./real-estate-service"]

