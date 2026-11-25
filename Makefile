.PHONY: build run format migrate-up migrate-down

build:
	@go build -o ./bin/api ./cmd/api

run: build
	@./bin/api

format:
	@go fmt ./...

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down