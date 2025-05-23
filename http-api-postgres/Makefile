.PHONY: start-db install start

# Start PostgreSQL database using Docker Compose
start-db:
	docker-compose up -d

# Install Go dependencies
install:
	go mod download
	go mod tidy

# Start the API server
start:
	go run cmd/api/main.go 