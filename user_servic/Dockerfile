FROM golang:1.22.2-alpine

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod ./
COPY go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь код
COPY . .

# Собираем приложение
RUN go mod tidy && go build -o /user_service ./cmd/server

EXPOSE 50051

CMD ["/user_service"]

