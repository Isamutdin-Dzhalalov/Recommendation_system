# Название исполняемого файла
BINARY_NAME=user_service

# Путь к основному пакету
MAIN_PACKAGE=./cmd/server

# Цель по умолчанию
all: build

# Сборка исполняемого файла
build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Запуск приложения
run: build
	./$(BINARY_NAME)

# Очистка созданных файлов
clean:
	rm -f $(BINARY_NAME)

# Установка зависимостей
deps:
	go mod tidy
	go mod download

