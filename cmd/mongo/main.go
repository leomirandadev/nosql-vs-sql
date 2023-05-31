package main

import (
	"context"
	"log"
	mongoRepo "mongo-vs-postgres/internal/repositories/mongo"
	"mongo-vs-postgres/internal/usecases/users"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Fatal("RUNNING MONGO GET")

	ctx := context.Background()
	mongoConn := newConnMongo().Database("testing")
	repo := mongoRepo.New(mongoConn)

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

func newConnMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("connection mongo")
	}

	return client
}
