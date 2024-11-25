# Используем официальный Go-образ в качестве основы
FROM golang:1.22-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum (для кеширования зависимостей)
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Строим приложение
RUN go build -o app .

# Используем минимальный образ для финальной версии
FROM alpine:latest

# Устанавливаем нужные зависимости для работы приложения (если нужно)
RUN apk --no-cache add ca-certificates

# Копируем приложение из стадии сборки
COPY --from=builder /app/app /app/app

# Указываем, что контейнер будет слушать на порту 8080
EXPOSE 8080

# Запускаем приложение
CMD ["/app/app"]
