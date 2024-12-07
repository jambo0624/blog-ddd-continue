.PHONY: all build test clean run migrate lint mock

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=blog-server
MAIN_PATH=cmd/main.go

# Build parameters
BUILD_DIR=build
BINARY_UNIX=$(BUILD_DIR)/$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BUILD_DIR)/$(BINARY_NAME).exe

all: test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

test:
	$(GOTEST) -v -race -cover ./internal/... ./tests/...

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

run:
	$(GOCMD) run $(MAIN_PATH)

# download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# generate mock files
mock:
	mockery --all --keeptree

# database init
db-init:
	DB_USER=$${DB_USER:-blog_user} \
	DB_PASSWORD=$${DB_PASSWORD:-your_default_password} \
	psql -U postgres -f database/init.sql 

# execute migrations
migrate:
	$(GOCMD) run cmd/migrate/main.go

migrate-test:
	GO_ENV=test $(GOCMD) run cmd/migrate/main.go

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
test-env:
	GO_ENV=test $(GOCMD) run $(MAIN_PATH)

# run in production environment
prod:
	GO_ENV=production $(GOCMD) run $(MAIN_PATH)

# Docker related commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME)

# help information
help:
	@echo "Make commands:"
	@echo "make build          - Build the application"
	@echo "make test           - Run tests"
	@echo "make clean          - Clean build files"
	@echo "make run            - Run the application"
	@echo "make deps           - Download and tidy dependencies"
	@echo "make db-init        - Initialize database"
	@echo "make mock           - Generate mock files"
	@echo "make migrate        - Run database migrations"
	@echo "make migrate-test   - Run database migrations in test environment"
	@echo "make lint           - Run linter"
	@echo "make dev            - Run in development mode"
	@echo "make test-env       - Run in test environment"
	@echo "make prod           - Run in production mode"
	@echo "make docker-build   - Build Docker image"
	@echo "make docker-run     - Run Docker container" 
