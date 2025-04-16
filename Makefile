.PHONY: build run test clean docker-build docker-up docker-down docker-logs help

# Variables
BINARY_NAME=main
DOCKER_COMPOSE=docker-compose
GO=go

# Colors for output
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m # No Color

help: ## Display this help message
	@echo "Available commands:"
	@echo ""
	@echo "$(GREEN)Development:$(NC)"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application locally"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo ""
	@echo "$(GREEN)Docker:$(NC)"
	@echo "  make docker-build - Build Docker images"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make docker-logs  - View Docker container logs"
	@echo ""

build: ## Build the application
	@echo "$(YELLOW)Building application...$(NC)"
	$(GO) build -o $(BINARY_NAME) ./cmd/app
	@echo "$(GREEN)Build completed!$(NC)"

run: ## Run the application locally
	@echo "$(YELLOW)Running application...$(NC)"
	$(GO) run ./cmd/app

test: ## Run tests
	@echo "$(YELLOW)Running tests...$(NC)"
	$(GO) test -v ./...

clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning...$(NC)"
	rm -f $(BINARY_NAME)
	@echo "$(GREEN)Clean completed!$(NC)"

docker-build: ## Build Docker images
	@echo "$(YELLOW)Building Docker images...$(NC)"
	$(DOCKER_COMPOSE) build
	@echo "$(GREEN)Docker build completed!$(NC)"

docker-up: ## Start Docker containers
	@echo "$(YELLOW)Starting Docker containers...$(NC)"
	$(DOCKER_COMPOSE) up -d
	@echo "$(GREEN)Docker containers started!$(NC)"

docker-down: ## Stop Docker containers
	@echo "$(YELLOW)Stopping Docker containers...$(NC)"
	$(DOCKER_COMPOSE) down
	@echo "$(GREEN)Docker containers stopped!$(NC)"

docker-logs: ## View Docker container logs
	@echo "$(YELLOW)Showing Docker container logs...$(NC)"
	$(DOCKER_COMPOSE) logs -f

# Default target
.DEFAULT_GOAL := help 