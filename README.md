# Docker Service API

A RESTful and gRPC API service for managing Docker containers with JWT authentication.

## Features

- User registration and authentication
- JWT-based authentication
- Docker container management:
  - List containers
  - Start containers
  - Stop containers
  - View container logs
- Dual API support:
  - REST API (port 8081)
  - gRPC API (port 8082)

## Prerequisites

- Go 1.23 or higher
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

The servers will start on:
- REST API: port 8081
- gRPC API: port 8082

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

The services will be available at:
- REST API: http://localhost:8081
- gRPC API: localhost:8082

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
make docker-build # Build Docker images
make docker-up   # Start Docker containers
make docker-down # Stop Docker containers
make clean       # Clean up all resources
```

## API Documentation

### REST API

#### Authentication

- POST /api/auth/register
  - Register a new user
  - Body: { "username": "string", "password": "string" }
  - Returns: { "token": "string" }

- POST /api/auth/login
  - Login with existing user
  - Body: { "username": "string", "password": "string" }
  - Returns: { "token": "string" }

#### Container Management

- GET /api/containers
  - List all containers
  - Requires: Authorization header with JWT token
  - Returns: Array of container objects

- POST /api/containers/{id}/start
  - Start a container
  - Requires: Authorization header with JWT token
  - Returns: { "message": "string" }

- POST /api/containers/{id}/stop
  - Stop a container
  - Requires: Authorization header with JWT token
  - Returns: { "message": "string" }

- GET /api/containers/{id}/logs
  - Get container logs
  - Requires: Authorization header with JWT token
  - Returns: { "logs": "string" }

### gRPC API

The gRPC API provides the same functionality as the REST API. The service definition can be found in `api/docker.proto`.

#### Authentication

- Register(stream AuthRequest) returns (stream AuthResponse)
- Login(stream AuthRequest) returns (stream AuthResponse)

#### Container Management

- ListContainers(stream Empty) returns (stream ContainerList)
- StartContainer(stream ContainerID) returns (stream OperationResponse)
- StopContainer(stream ContainerID) returns (stream OperationResponse)
- GetContainerLogs(stream ContainerID) returns (stream ContainerLogs)

## Development

### Project Structure

```
.
├── api/                    # API definitions
│   ├── docker.proto       # gRPC service definition
│   └── docker_grpc.pb.go  # Generated gRPC code
├── cmd/                   # Application entry points
│   └── app/              # Main application
├── internal/             # Internal packages
│   ├── config/          # Configuration
│   ├── core/            # Core business logic
│   ├── models/          # Database models
│   └── transport/       # API transport
│       ├── GRPC/        # gRPC server
│       └── REST/        # REST server
├── migrations/          # Database migrations
└── tests/              # Test files
```

### Adding New Features

1. Update the API definition in `api/docker.proto` for gRPC
2. Implement the new feature in the core package
3. Add the feature to both REST and gRPC servers
4. Update documentation

## License

MIT

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
