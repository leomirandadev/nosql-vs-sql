package main

import (
	"log"
	"nosql-vs-sql/configs"
	"nosql-vs-sql/internal/exec"
	mongoRepo "nosql-vs-sql/internal/repositories/mongo"
	"nosql-vs-sql/internal/usecases"
)

func main() {
	log.Println("RUNNING MONGO GET")

	var mongoConn = configs.NewConnMongo().Database("mongovspostgres")
	var repo = mongoRepo.New(mongoConn)
	var useCases = usecases.New(repo)

	for i := 0; i < 3; i++ {
		exec.Run(useCases, "6477c2ffc1adaf0e33063c09", "6477c297245a0416d6a64bcb")
	}
}
