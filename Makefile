.PHONY: build up down clean

# Variables
DOCKER_COMPOSE = docker compose
BINARY_NAME = main
GO_CMD = go

# Colors for output
GREEN = \033[0;32m
RED = \033[0;31m
NC = \033[0m

# Build service
build:
	@echo "$(GREEN)Building service...$(NC)"
	$(DOCKER_COMPOSE) build app

# Run service
up: build
	@echo "$(GREEN)Starting service...$(NC)"
	$(DOCKER_COMPOSE) up -d app postgres

# Stop service
down:
	@echo "$(GREEN)Stopping service...$(NC)"
	$(DOCKER_COMPOSE) down

# Clean up
clean:
	@echo "$(GREEN)Cleaning up...$(NC)"
	$(DOCKER_COMPOSE) down -v

# Help command
help:
	@echo "Available commands:"
	@echo "  make build    - Build service"
	@echo "  make up      - Start service"
	@echo "  make down    - Stop service"
	@echo "  make clean   - Clean up all resources" 