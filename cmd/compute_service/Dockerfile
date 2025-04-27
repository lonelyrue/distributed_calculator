# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

# Копируем только go.mod и go.sum сначала
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Теперь копируем весь проект
COPY . .

# Ставим рабочую директорию
WORKDIR /app/cmd/compute_service

# Важно: включаем режим модулей
ENV GO111MODULE=on

# Собираем бинарник
RUN go build -mod=mod -o compute_service main.go

# Команда по умолчанию
CMD ["./compute_service"]