PHONY := test lint run

all: test lint run

test:
	@echo "Running Go tests..."
	go test ./...

lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

run:
	@echo "Running example..."
	go run example/main.go