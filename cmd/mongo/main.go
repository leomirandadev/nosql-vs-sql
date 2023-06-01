package main

import (
	"context"
	"log"
	"nosql-vs-sql/configs"
	mongoRepo "nosql-vs-sql/internal/repositories/mongo"
	"nosql-vs-sql/internal/usecases/users"
	"time"
)

func main() {
	log.Println("RUNNING MONGO GET")

	for i := 0; i < 3; i++ {
		exec()
	}
}

var ctx = context.Background()
var mongoConn = configs.NewConnMongo().Database("mongovspostgres")
var repo = mongoRepo.New(mongoConn)
var usersUseCase = users.New(repo)

func exec() {
	now := time.Now()
	users, err := usersUseCase.GetDetailsAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), "qty users:", len(users))

	now = time.Now()
	user, err := usersUseCase.GetDetailsByID(ctx, "6477c2ffc1adaf0e33063c09")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), user)
}
