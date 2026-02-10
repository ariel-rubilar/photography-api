# Makefile for photography-api
# Basic targets: build, run, install, test, fmt, vet, tidy, mod-download, clean, help

APP_NAME := photography-api
CMD_PATH := ./cmd/api
BIN_DIR := bin
BINARY := $(BIN_DIR)/$(APP_NAME)

.PHONY: all build run start install fmt vet tidy mod-download test clean setup precommit help

all: build

# Build the binary to ./bin/photography-api
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BINARY) $(CMD_PATH)
	@echo "Built $(BINARY)"

# Run the application (compiles and runs in module mode)
run:
	@echo "Running $(APP_NAME) via 'go run'..."
	@doppler run  -- go run $(CMD_PATH)

# Alias for run (semantic)
start: run

# Install the binary into the toolchain (GOBIN/GOPATH/bin)
install:
	@echo "Installing $(APP_NAME)..."
	@go install $(CMD_PATH)

# Formatting and lint helpers
fmt:
	@echo "Running go fmt..."
	@go fmt ./...

vet:
	@echo "Running go vet..."
	@go vet ./...

# Go module helpers
tidy:
	@echo "Running go mod tidy..."
	@go mod tidy

mod-download:
	@echo "Downloading modules..."
	@go mod download

# Tests
test:
	@echo "Running tests..."
	@go test ./...

# Remove built artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)


# Setup Lefthook for git hooks
setup:
	@echo "🔧 Installing Lefthook..."
	@lefthook install
	@echo "✅ Lefthook installed"
	@doppler setup --project photography-api --config dev_personal

precommit:
	@echo "Running pre-commit hooks..."
	@lefthook run pre-commit

# Help
help:
	@printf "Available targets:\n"
	@printf "  make build         - Build binary into ./bin\n"
	@printf "  make run|start     - Run the app (go run ./cmd/api)\n"
	@printf "  make install       - Install binary to GOBIN/GOPATH/bin\n"
	@printf "  make fmt           - Run go fmt on the repo\n"
	@printf "  make vet           - Run go vet on the repo\n"
	@printf "  make tidy          - Run go mod tidy\n"
	@printf "  make mod-download  - Download modules\n"
	@printf "  make test          - Run tests\n"
	@printf "  make clean         - Remove ./bin\n"
	@printf "  make setup         - Setup Lefthook for git hooks\n"
	@printf "  make precommit     - Run pre-commit hooks via Lefthook\n"
