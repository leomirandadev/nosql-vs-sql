package main

import (
	"context"
	"log"
	"mongo-vs-postgres/configs"
	"mongo-vs-postgres/internal/repositories/postgres"
	"mongo-vs-postgres/internal/usecases/users"
	"time"
)

func main() {
	log.Println("RUNNING POSTGRES GET")

	ctx := context.Background()
	postgressConn := configs.NewPostgresConn()
	repo := postgres.New(postgressConn)

	usersUseCase := users.New(repo)

	now := time.Now()
	users, err := usersUseCase.GetAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), "qty users:", len(users))

	now = time.Now()
	user, err := usersUseCase.GetByID(ctx, "1")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), user)
}
