package main

import (
	"context"
	"log"
	"mongo-vs-postgres/configs"
	"mongo-vs-postgres/internal/repositories/sql"
	"mongo-vs-postgres/internal/usecases/users"
	"time"
)

func main() {
	log.Println("RUNNING POSTGRES GET")

	for i := 0; i < 3; i++ {
		exec()
	}
}

var ctx = context.Background()
var mariadbConn = configs.NewMariaDBConn()
var repo = sql.New(mariadbConn)
var usersUseCase = users.New(repo)

func exec() {
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
