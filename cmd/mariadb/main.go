package main

import (
	"log"
	"nosql-vs-sql/configs"
	"nosql-vs-sql/internal/exec"
	"nosql-vs-sql/internal/repositories/sql"
	"nosql-vs-sql/internal/usecases"
)

func main() {
	log.Println("RUNNING MARIADB GET")

	var mariadbConn = configs.NewMariaDBConn()
	var repo = sql.New(mariadbConn, false)
	var useCases = usecases.New(repo)

	for i := 0; i < 3; i++ {
		exec.Run(useCases, "1", "1")
	}

}
