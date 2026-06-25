

PATH_MIGRATION := ./internal/database/migrations

create-migration:
	@echo "Creating migration file..."
	@goose -dir $(PATH_MIGRATION) create alterar_nome sql

run-app:
	@echo "Running the application..."
	@go run cmd/server/main.go

run-linter:
	@echo "Running linter..."
	@golangci-lint run