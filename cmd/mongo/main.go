package main

import (
	"context"
	"log"
	"mongo-vs-postgres/configs"
	mongoRepo "mongo-vs-postgres/internal/repositories/mongo"
	"mongo-vs-postgres/internal/usecases/users"
	"time"
)

func main() {
	log.Println("RUNNING MONGO GET")

	ctx := context.Background()
	mongoConn := configs.NewConnMongo().Database("mongovspostgres")
	repo := mongoRepo.New(mongoConn)

	usersUseCase := users.New(repo)

	now := time.Now()
	users, err := usersUseCase.GetAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), "qty users:", len(users))

	now = time.Now()
	user, err := usersUseCase.GetByID(ctx, "6477c2ffc1adaf0e33063c09")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), user)
}
