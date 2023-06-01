package main

import (
	"log"
	"nosql-vs-sql/configs"
	"nosql-vs-sql/internal/exec"
	"nosql-vs-sql/internal/repositories/sql"
	"nosql-vs-sql/internal/usecases"
)

func main() {
	log.Println("RUNNING POSTGRES GET")

	var postgressConn = configs.NewPostgresConn()
	var repo = sql.New(postgressConn, true)
	var useCases = usecases.New(repo)

	for i := 0; i < 3; i++ {
		exec.Run(useCases, "1", "1")
	}
}
