# Build stage
FROM golang:1.23-alpine AS builder

# Install git and build dependencies
RUN apk add --no-cache git build-base

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main cmd/app/main.go

# Final stage
FROM golang:1.23-alpine

WORKDIR /app

# Install Docker client
RUN apk add --no-cache docker-cli

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy test files
COPY --from=builder /app/go.mod /app/go.sum ./
COPY --from=builder /app/internal ./internal

# Set environment variables
ENV TEST_MODE=false
ENV GIN_MODE=release

# Expose port
EXPOSE 8081 8082

# Set environment variables
ENV JWT_SECRET_KEY=your-secret-key-here
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=docker_service

# Command to run the application
CMD ["./main"] 