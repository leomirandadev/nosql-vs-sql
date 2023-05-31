DB_CONNECTION = "user=root password=root dbname=mongovspostgres host=localhost port=5432 sslmode=disable"

install:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

db-up:
	@docker compose -f "docker/docker-compose.yml" up -d --build

db-down:
	@docker compose -f "docker/docker-compose.yml" down

mig-status-postgres: 
	@goose postgres $(DB_CONNECTION) status

mig-create-postgres: 
	@goose -dir ./migrations/postgres create $(MIG_NAME) sql 

mig-up-postgres: 
	@goose -dir ./migrations/postgres postgres $(DB_CONNECTION) up

mig-down-postgres: 
	@goose -dir ./migrations/postgres postgres $(DB_CONNECTION) down

mig-status: 
	@make mig-status-postgres

mig-up: 
	@make mig-up-postgres

mig-down: 
	@make mig-down-postgres

run-postgres:
	@go run cmd/postgres/*.go

run-mongo:
	@go run cmd/mongo/*.go