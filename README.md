HTTP-сервис с Redis, HMAC-SHA512, PostgreSQL.

## Требования

- Go 1.24+
- Docker

## Быстрый старт (Docker)
```bash
git clone https://github.com/Boombstick/firstprogram.git
cd firstprogram
```

Создай файл `.env.docker`:
```env
SERVER_PORT=8080
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=default
POSTGRES_PASSWORD=password
POSTGRES_DB=test
REDIS_HOST=redis
REDIS_PORT=6379
```

Запусти:
```bash
docker-compose up --build
```

Сервис доступен на http://localhost:8080
Swagger UI: http://localhost:8080/swagger/index.html

## Локальная разработка

Создай файл `.env`:
```env
SERVER_PORT=8080
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=default
POSTGRES_PASSWORD=password
POSTGRES_DB=test
REDIS_HOST=localhost
REDIS_PORT=6379
```
```bash
swag init
go mod tidy
go run main.go
```

## Тесты
```bash
go test ./services/ -v
```
