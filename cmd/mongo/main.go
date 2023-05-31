package main

import (
	"context"
	"log"
	"mongo-vs-postgres/configs"
	mongoRepo "mongo-vs-postgres/internal/repositories/mongo"
	"mongo-vs-postgres/internal/usecases/users"
)

func main() {
	log.Println("RUNNING MONGO GET")

	ctx := context.Background()
	mongoConn := configs.NewConnMongo().Database("mongovspostgres")
	repo := mongoRepo.New(mongoConn)

	usersUseCase := users.New(repo)

	users, err := usersUseCase.GetAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println("qty users:", len(users))

	user, err := usersUseCase.GetByID(ctx, "6477a1dc939fc4752dc8d085")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(user)
}
