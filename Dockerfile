# Используем официальный образ Go в качестве builder
FROM golang:1.22-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем только go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Строим приложение
RUN go build -o /app/cmd/main cmd/main.go

# Используем минимальный образ для финальной версии
FROM alpine:latest

# Устанавливаем нужные зависимости для работы приложения (если нужно)
#RUN apk --no-cache add ca-certificates

# Копируем скомпилированный бинарник из builder стадии
COPY --from=builder /app/cmd/main /app/main

# Открываем порт, на котором будет слушать приложение
EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]
