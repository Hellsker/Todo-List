# Используем официальное изображение Golang как базовое
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в контейнер
COPY . .

# Компилируем приложение
RUN go build -o ./cmd/app/main ./cmd/app

# Указываем команду для запуска приложения
CMD ["./cmd/app/main"]

