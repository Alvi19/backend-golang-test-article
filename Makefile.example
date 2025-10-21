APP_NAME=Backend-Golang-Test
APP_PORT=8080
APP_ENV=development
JWT_SECRET=supersecretkey123 

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=article_db
DB_SSLMODE=disable

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# ---------- Path ----------
MIGRATIONS_DIR=./migrations

# ---------- Command ----------
MIGRATE_CMD=migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)"

## ---------- Commands ---------- ##

# Jalankan semua migration (up)
migrate-up:
	$(MIGRATE_CMD) up

# Rollback 1 langkah (down 1)
migrate-down:
	$(MIGRATE_CMD) down 1

# Reset semua migration
migrate-reset:
	$(MIGRATE_CMD) drop -f
	$(MIGRATE_CMD) up

# Buat migration baru
# Usage: make migrate-create name=create_users_table
migrate-create:
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Run app (langsung pake variabel Makefile)
run:
	APP_NAME=$(APP_NAME) \
	APP_ENV=$(APP_ENV) \
	APP_PORT=$(APP_PORT) \
	DB_HOST=$(DB_HOST) \
	DB_PORT=$(DB_PORT) \
	DB_USER=$(DB_USER) \
	DB_PASSWORD=$(DB_PASSWORD) \
	DB_NAME=$(DB_NAME) \
	DB_SSLMODE=$(DB_SSLMODE) \
	JWT_SECRET=$(JWT_SECRET) \
	go run ./cmd/server

# Jalankan migration + server
run-migrate: migrate-up run

# Build binary
build:
	go build -o bin/tas-api ./cmd/server

# Tidy dependencies
tidy:
	go mod tidy

# Generate swagger docs
swagger-gen:
	swag init -g cmd/server/main.go -o docs