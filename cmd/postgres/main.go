package main

import (
	"context"
	"log"
	"mongo-vs-postgres/configs"
	"mongo-vs-postgres/internal/repositories/postgres"
	"mongo-vs-postgres/internal/usecases/users"
)

func main() {
	log.Println("RUNNING POSTGRES GET")

	ctx := context.Background()
	postgressConn := configs.NewPostgresConn()
	repo := postgres.New(postgressConn)

	usersUseCase := users.New(repo)

	users, err := usersUseCase.GetAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println("qty users:", len(users))

	user, err := usersUseCase.GetByID(ctx, "1")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(user)
}
