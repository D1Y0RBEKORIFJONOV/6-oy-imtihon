FROM golang:1.23.0 AS builder

WORKDIR /app

# Устанавливаем переменную CGO_ENABLED=0 для сборки статического бинарного файла
ENV CGO_ENABLED=0 GOOS=linux

COPY . .

# Собираем бинарник
RUN go build -o main cmd/app/main.go

# Используем alpine для минимального размера образа
FROM alpine:latest

# Устанавливаем необходимые сертификаты
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем сгенерированный бинарник
COPY --from=builder /app/main .

# Копируем конфигурационные файлы в корень контейнера
COPY auth.conf /app/auth.conf
COPY auth.csv /app/auth.csv

# Копируем сертификаты
COPY internal/config/certf/localhost.pem /localhost.pem
COPY internal/config/certf/localhost-key.pem /localhost-key.pem

# Даем права на сертификаты
RUN chmod 600 /localhost.pem /localhost-key.pem

# Устанавливаем точку входа для запуска приложения
CMD ["/app/main"]
