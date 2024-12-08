.PHONY: all build run-test clean run deps init-db migrate migrate-test lint build-linux build-windows dev test prod

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=blog-server
MAIN_PATH=cmd/main.go
MIGRATE_PATH=cmd/migrate/main.go

# Build parameters
BUILD_DIR=bin
BINARY_UNIX=$(BUILD_DIR)/$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BUILD_DIR)/$(BINARY_NAME).exe

all: run-test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

run-test:
	$(GOTEST) -v -race -cover ./tests/...

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

run:
	$(GOCMD) run $(MAIN_PATH)

# download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# database init
DB_USER ?= blog_user
DB_PASSWORD ?= your_default_password

init-db:
	DB_USER=$${DB_USER} \
	DB_PASSWORD=$${DB_PASSWORD} \
	psql -U postgres -f database/init.sql

# execute migrations
migrate:
	$(GOCMD) run $(MIGRATE_PATH)

migrate-test:
	GO_ENV=test $(GOCMD) run $(MIGRATE_PATH)

# run code lint
lint:
	golangci-lint run

# cross compile
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) $(MAIN_PATH)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WINDOWS) $(MAIN_PATH)

# run in development mode
dev:
	GO_ENV=development $(GOCMD) run $(MAIN_PATH)

# run in test environment
test:
	GO_ENV=test $(GOCMD) run $(MAIN_PATH)

# run in production environment
prod:
	GO_ENV=production $(GOCMD) run $(MAIN_PATH)

# help information
help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  all          	Build and run tests"
	@echo "  build        	Build the application"
	@echo "  run-test     	Run tests"
	@echo "  clean        	Clean the build directory"
	@echo "  run          	Run the application, default is development environment"
	@echo "  deps         	Download dependencies"
	@echo "  init-db      	Initialize the database"
	@echo "  migrate      	Execute migrations"
	@echo "  migrate-test		Execute migrations in test environment"
	@echo "  lint         	Run code lint"
	@echo "  build-linux  	Build the application for Linux"
	@echo "  build-windows	Build the application for Windows"
	@echo "  dev          	Run the application in development environment"
	@echo "  test         	Run the application in test environment"
	@echo "  prod         	Run the application in production environment"
