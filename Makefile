# Default env file
ENV_FILE=.env.local

# Binary name
BINARY_NAME=app

# Build binary
build:
	go build -o $(BINARY_NAME) ./cmd

# Run app locally
run:
	go run ./cmd/main.go

# Test all packages
test:
	go test ./... -v

# Run app with env
run-env:
	ENV_FILE=$(ENV_FILE) go run ./cmd/main.go

# Load env then run tests
test-env:
	ENV_FILE=$(ENV_FILE) go test ./... -v

# Docker
docker-up:
	docker compose up --build

docker-down:
	docker compose down

# Clean Docker data
docker-clean:
	docker compose down -v --remove-orphans

# Help
help:
	@echo "make build           - Build Go binary"
	@echo "make run             - Run app locally"
	@echo "make run-env         - Run with .env.local"
	@echo "make test            - Run tests"
	@echo "make lint            - Run linter"
	@echo "make docker-up       - Build & run Docker"
	@echo "make docker-down     - Stop Docker"
	@echo "make docker-clean    - Full cleanup of containers & volumes"
