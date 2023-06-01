package states

import (
	"context"
	"nosql-vs-sql/configs"
	mongoRepo "nosql-vs-sql/internal/repositories/mongo"
	"nosql-vs-sql/internal/repositories/sql"
	"testing"

	_ "github.com/lib/pq"
)

var ctx = context.Background()

// mongo
var mongoConn = configs.NewConnMongo().Database("mongovspostgres")
var repoMongo = mongoRepo.New(mongoConn)
var statesUseCaseMongo = New(repoMongo)

// postgres
var postgressConn = configs.NewPostgresConn()
var repoPostgres = sql.New(postgressConn, true)
var statesUseCasePostgres = New(repoPostgres)

// mariadb
var mariadbConn = configs.NewPostgresConn()
var repoMariaDB = sql.New(mariadbConn, false)
var statesUseCaseMariaDB = New(repoMariaDB)

func BenchmarkStatesGetAllPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		statesUseCasePostgres.GetAll(ctx)
	}
}

func BenchmarkStatesGetAllMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		statesUseCaseMariaDB.GetAll(ctx)
	}
}

func BenchmarkStatesGetAllMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		statesUseCaseMongo.GetAll(ctx)
	}
}

func BenchmarkStatesGetByIDPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		statesUseCasePostgres.GetByID(ctx, "1")
	}
}

func BenchmarkStatesGetByIDMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		statesUseCaseMariaDB.GetByID(ctx, "1")
	}
}

func BenchmarkStatesGetByIDMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		statesUseCaseMongo.GetByID(ctx, "6477c297245a0416d6a64bcb")
	}
}
