.PHONY: build run format format-sql migrate-up migrate-down

build:
	@go build -o ./bin/api ./cmd/api

run: build
	@./bin/api

format:
	@go fmt ./...
	@$(MAKE) format-sql

format-sql:
	@pg_format --inplace migrations/*.sql

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down