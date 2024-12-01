# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o main cmd/main.go

# Run the application
run:
	@go run cmd/main.go



# Test the application
test:
	@echo "Testing..."
	@go test ./... -v



# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: all build run test clean
