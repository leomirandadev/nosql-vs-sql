DB_CONNECTION_POSTGRES = "user=root password=root dbname=mongovspostgres host=localhost port=5432 sslmode=disable"
DB_CONNECTION_MARIADB = "root:root@(localhost:3306)/mongovspostgres"
install:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

db-up:
	@docker compose -f "docker/docker-compose.yml" up -d --build

db-down:
	@docker compose -f "docker/docker-compose.yml" down

mig-status-postgres: 
	@goose postgres $(DB_CONNECTION_POSTGRES) status

mig-create-postgres: 
	@goose -dir ./migrations/postgres create $(MIG_NAME) sql 

mig-up-postgres: 
	@goose -dir ./migrations/postgres postgres $(DB_CONNECTION_POSTGRES) up

mig-down-postgres: 
	@goose -dir ./migrations/postgres postgres $(DB_CONNECTION_POSTGRES) down

mig-status-mariadb: 
	@goose mysql $(DB_CONNECTION_MARIADB) status

mig-create-mariadb: 
	@goose -dir ./migrations/mariadb create $(MIG_NAME) sql 

mig-up-mariadb: 
	@goose -dir ./migrations/mariadb mysql $(DB_CONNECTION_MARIADB) up

mig-down-mariadb: 
	@goose -dir ./migrations/mariadb mysql $(DB_CONNECTION_MARIADB) down

mig-status: 
	@make mig-status-postgres
	@make mig-up-mariadb

mig-up: 
	@make mig-up-postgres
	@make mig-up-mariadb

mig-down: 
	@make mig-down-postgres
	@make mig-up-mariadb

run-postgres:
	@go run cmd/postgres/*.go

run-mariadb:
	@go run cmd/mariadb/*.go

run-mongo:
	@go run cmd/mongo/*.go

run-bench:
	@go test -v ./... -bench=. -benchmem