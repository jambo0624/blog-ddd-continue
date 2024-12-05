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

# 依赖管理
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# 生成 mock
mock:
	mockery --all --keeptree

# 数据库迁移
migrate:
	$(GOCMD) run cmd/migrate/main.go

# 代码检查
lint:
	golangci-lint run

# 交叉编译
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) $(MAIN_PATH)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WINDOWS) $(MAIN_PATH)

# 开发模式运行
dev:
	GO_ENV=development $(GOCMD) run $(MAIN_PATH)

# 测试环境运行
test-env:
	GO_ENV=test $(GOCMD) run $(MAIN_PATH)

# 生产环境运行
prod:
	GO_ENV=production $(GOCMD) run $(MAIN_PATH)

# Docker 相关命令
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME)

# 帮助信息
help:
	@echo "Make commands:"
	@echo "make build          - Build the application"
	@echo "make test           - Run tests"
	@echo "make clean          - Clean build files"
	@echo "make run           - Run the application"
	@echo "make deps          - Download and tidy dependencies"
	@echo "make mock          - Generate mock files"
	@echo "make migrate       - Run database migrations"
	@echo "make lint          - Run linter"
	@echo "make dev           - Run in development mode"
	@echo "make test-env      - Run in test environment"
	@echo "make prod          - Run in production mode"
	@echo "make docker-build  - Build Docker image"
	@echo "make docker-run    - Run Docker container" 