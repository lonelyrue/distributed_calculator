# syntax=docker/dockerfile:1
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/calc_service

RUN go build -o calc_service main.go

EXPOSE 8080

CMD ["./calc_service"]