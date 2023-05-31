package main

import (
	"context"
	"database/sql"
	"log"
	"mongo-vs-postgres/internal/repositories/postgres"
	"mongo-vs-postgres/internal/usecases/users"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("RUNNING POSTGRES GET")

	ctx := context.Background()
	postgressConn := newPostgresConn()
	repo := postgres.New(postgressConn)

	usersUseCase := users.New(repo)

	users, err := usersUseCase.GetAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(users)

	user, err := usersUseCase.GetByID(ctx, "1")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(user)
}

func newPostgresConn() *sql.DB {
	conn, err := sql.Open("postgres", "user=root password=root dbname=mongovspostgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("[ERROR]", err)
	}
	return conn
}
