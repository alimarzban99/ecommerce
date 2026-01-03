.PHONY: migrate-create migrate-up migrate-down migrate-down-all migrate-force migrate-version migrate-status migrate-drop

# Database connection string (can be overridden with environment variables)
# Note: CLI-based commands may not work if migrate CLI doesn't have postgres driver
DB_URL ?= postgres://user:pass@localhost:5432/ecommerce?sslmode=disable

# Migration directory
MIGRATIONS_DIR = database/migrations

# Create a new migration file
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=create_users_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Run all pending migrations (uses Go, reads from .env)
migrate-up:
	@go run cmd/migrate/main.go up

# Rollback the last migration (uses Go, reads from .env)
migrate-down:
	@go run cmd/migrate/main.go down

# Rollback all migrations (uses Go, reads from .env)
migrate-down-all:
	@go run cmd/migrate/main.go down -all

# Force migration version (uses Go, reads from .env)
migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Usage: make migrate-force version=1"; \
		exit 1; \
	fi
	@go run cmd/migrate/main.go force $(version)

# Show current migration version (uses Go, reads from .env)
migrate-version:
	@go run cmd/migrate/main.go version

# Show detailed migration status (uses Go, reads from .env)
migrate-status:
	@go run cmd/migrate/main.go status

# CLI-based commands (may not work if migrate CLI doesn't have postgres driver)
# Use these only if you have migrate CLI with postgres support installed
migrate-up-cli:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down-cli:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-version-cli:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

