all: build test

.PHONY: build
build:
	@echo "Building..."
	@go build ./...

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -cover -parallel 1 ./...

.PHONY: format
format:
	@echo "Formatting code..."
	@go fmt ./...