# Используем официальный образ Golang
FROM golang:1.22.2-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем оставшиеся файлы в контейнер
COPY . .

# Собираем приложение
RUN go build -o /recommendation_service ./cmd/server

# Указываем порт, на котором приложение будет слушать
EXPOSE 50052

# Команда для запуска приложения
CMD [ "/recommendation_service" ]

