# Simple Makefile for a Go project
# Load environment variables from .env
include .env
# Variables from .env will be available here
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=$(CONN_STRING)
GOOSE_MIGRATION_DIR=internal/migration
all: build

build:
	@echo "Building..."
	
	
	@go build -o main.exe cmd/server/main.go

# Run the application
run:
	@go run cmd/server/main.go


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
	@go test ./... -v


# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload

watch:
	@air

up:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) up

down:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) down

create-migration:
	cd $(GOOSE_MIGRATION_DIR) && goose create $(name) sql
	
sqlc-gen:	
	@sqlc generate

wire:
	cd internal/wire && wire

.PHONY: all build run test clean watch up down create-migration