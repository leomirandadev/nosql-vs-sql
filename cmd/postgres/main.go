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
	log.Fatal("RUNNING POSTGRES GET")

	ctx := context.Background()
	postgressConn := newPostgresConn()
	repo := postgres.New(postgressConn)

	usersUseCase := users.New(repo)

	users, err := usersUseCase.GetAll(ctx)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	log.Println(users)

	user, err := usersUseCase.GetByID(ctx, "1")
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	log.Println(user)
}

func newPostgresConn() *sql.DB {
	conn, err := sql.Open("postgres", "root:root@tcp(127.0.0.1:3306)/your_database_name")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
