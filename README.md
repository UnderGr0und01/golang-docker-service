# Docker Service API

A RESTful API service for managing Docker containers with JWT authentication.

## Features

- User registration and authentication
- JWT-based authentication
- Docker container management:
  - List containers
  - Start containers
  - Stop containers
  - View container logs

## Prerequisites

- Go 1.16 or higher
- Docker
- Docker Compose
- PostgreSQL (for user authentication)
- Make (optional, but recommended)

## Installation

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/docker-service.git
cd docker-service
```

2. Set up environment variables:
```bash
export JWT_SECRET_KEY=your-secret-key-here
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=your_db_user
export DB_PASSWORD=your_db_password
export DB_NAME=your_db_name
```

3. Run the application:
```bash
# Using make
make run

# Or directly
go run cmd/app/main.go
```

The server will start on port 8081.

### Docker Deployment

1. Clone the repository:
```bash
git clone https://github.com/yourusername/docker-service.git
cd docker-service
```

2. Build and start the containers:
```bash
# Using make
make docker-build
make docker-up

# Or using docker-compose directly
docker-compose up --build
```

The service will be available at http://localhost:8081

To run in detached mode:
```bash
docker-compose up -d --build
```

To stop the services:
```bash
# Using make
make docker-down

# Or using docker-compose directly
docker-compose down
```

To stop and remove volumes (including database data):
```bash
docker-compose down -v
```

### Make Commands

The project includes a Makefile with the following commands:

```bash
make help        # Show all available commands
make build       # Build the application
make run         # Run the application locally
make test        # Run tests
make clean       # Clean build artifacts
make docker-build # Build Docker images
make docker-up   # Start Docker containers
make docker-down # Stop Docker containers
make docker-logs # View Docker container logs
```

## API Documentation

### Authentication

#### Register a new user
```http
POST /register
Content-Type: application/json

{
    "username": "your_username",
    "password": "your_password"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
    "username": "your_username",
    "password": "your_password"
}
```

Response will include a JWT token:
```json
{
    "token": "your.jwt.token"
}
```

### Container Management

All container management endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer your.jwt.token
```

#### List Containers
```http
GET /api/containers/
```

#### Start Container
```http
POST /api/start/:id
```

#### Stop Container
```http
POST /api/stop/:id
```

#### Get Container Logs
```http
GET /api/logs/:id
```

## Security

- All passwords are hashed using bcrypt
- JWT tokens are used for authentication
- Protected routes require valid JWT token
- No sensitive data is logged

---

# Docker Service API (Русский)

RESTful API сервис для управления Docker контейнерами с JWT аутентификацией.

## Возможности

- Регистрация и аутентификация пользователей
- Аутентификация на основе JWT
- Управление Docker контейнерами:
  - Список контейнеров
  - Запуск контейнеров
  - Остановка контейнеров
  - Просмотр логов контейнеров

## Требования

- Go 1.16 или выше
- Docker
- Docker Compose
- PostgreSQL (для аутентификации пользователей)
- Make (опционально, но рекомендуется)

## Установка

### Локальная разработка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/docker-service.git
cd docker-service
```

2. Настройте переменные окружения:
```bash
export JWT_SECRET_KEY=ваш-секретный-ключ
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=ваш_пользователь_бд
export DB_PASSWORD=ваш_пароль_бд
export DB_NAME=ваше_имя_бд
```

3. Запустите приложение:
```bash
# Используя make
make run

# Или напрямую
go run cmd/app/main.go
```

Сервер запустится на порту 8081.

### Запуск в Docker

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/docker-service.git
cd docker-service
```

2. Соберите и запустите контейнеры:
```bash
# Используя make
make docker-build
make docker-up

# Или используя docker-compose напрямую
docker-compose up --build
```

Сервис будет доступен по адресу http://localhost:8081

Для запуска в фоновом режиме:
```bash
docker-compose up -d --build
```

Для остановки сервисов:
```bash
# Используя make
make docker-down

# Или используя docker-compose напрямую
docker-compose down
```

Для остановки и удаления томов (включая данные базы данных):
```bash
docker-compose down -v
```

### Команды Make

Проект включает Makefile со следующими командами:

```bash
make help        # Показать все доступные команды
make build       # Собрать приложение
make run         # Запустить приложение локально
make test        # Запустить тесты
make clean       # Очистить артефакты сборки
make docker-build # Собрать Docker образы
make docker-up   # Запустить Docker контейнеры
make docker-down # Остановить Docker контейнеры
make docker-logs # Просмотреть логи Docker контейнеров
```

## Документация API

### Аутентификация

#### Регистрация нового пользователя
```http
POST /register
Content-Type: application/json

{
    "username": "ваше_имя_пользователя",
    "password": "ваш_пароль"
}
```

#### Вход
```http
POST /login
Content-Type: application/json

{
    "username": "ваше_имя_пользователя",
    "password": "ваш_пароль"
}
```

Ответ будет содержать JWT токен:
```json
{
    "token": "ваш.jwt.токен"
}
```

### Управление контейнерами

Все эндпоинты управления контейнерами требуют JWT аутентификации. Включите токен в заголовок Authorization:
```
Authorization: Bearer ваш.jwt.токен
```

#### Список контейнеров
```http
GET /api/containers/
```

#### Запуск контейнера
```http
POST /api/start/:id
```

#### Остановка контейнера
```http
POST /api/stop/:id
```

#### Получение логов контейнера
```http
GET /api/logs/:id
```

## Безопасность

- Все пароли хешируются с использованием bcrypt
- Для аутентификации используются JWT токены
- Защищенные маршруты требуют действительный JWT токен
- Чувствительные данные не логируются
