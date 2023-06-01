package main

import (
	"context"
	"log"
	"nosql-vs-sql/configs"
	"nosql-vs-sql/internal/repositories/sql"
	"nosql-vs-sql/internal/usecases/users"
	"time"
)

func main() {
	log.Println("RUNNING POSTGRES GET")

	for i := 0; i < 3; i++ {
		exec()
	}
}

var ctx = context.Background()
var postgressConn = configs.NewPostgresConn()
var repo = sql.New(postgressConn, true)
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
	user, err := usersUseCase.GetDetailsByID(ctx, "1")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), user)
}
