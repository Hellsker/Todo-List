# Todo List API

Простое API Todo List для создание задач на день.

### Стэк:
- База данных: PostgreSQL
- Веб-фреймворк: Gin
- Драйвер БД: pgx
- Парсер конфига: Clean Env
- Логгер : slog
- Swagger
- Контейнеризация: Docker Compose

### Запуск:
``` 
docker-compose up
```
### Swagger
``` 
make swag init
```
### Migration
``` 
make migrate-up
make migrate-down
```
### Go run
``` 
make start
```
### Тесты
Не успел написать, не хватило времени.