# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o main.exe cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	air;

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./internal/database/migrations ${name}

## db/migrations/up: apply all database up migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running Up migrations...'

	migrate -path=./internal/database/migrations -database postgres://postgres:root@localhost:5432/motorq?sslmode=disable up

	
## db/migrations/down: apply all database down migrations
.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running Down migrations...'
	migrate -path=./internal/database/migrations -database postgres://postgres:root@localhost:5432/motorq?sslmode=disable down

.PHONY: db/migrations/reload
db/migrations/reload: db/migrations/down db/migrations/up


.PHONY: all build run test clean