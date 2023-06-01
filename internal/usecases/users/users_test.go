package users

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
var usersUseCaseMongo = New(repoMongo)

// postgres
var postgressConn = configs.NewPostgresConn()
var repoPostgres = sql.New(postgressConn, true)
var usersUseCasePostgres = New(repoPostgres)

// mariadb
var mariadbConn = configs.NewPostgresConn()
var repoMariaDB = sql.New(mariadbConn, false)
var usersUseCaseMariaDB = New(repoMariaDB)

func BenchmarkGetDetailsAllPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetDetailsAll(ctx)
	}
}

func BenchmarkGetDetailsAllMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMariaDB.GetDetailsAll(ctx)
	}
}

func BenchmarkGetDetailsAllMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetDetailsAll(ctx)
	}
}

func BenchmarkGetDetailsByIDPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetDetailsByID(ctx, "1")
	}
}

func BenchmarkGetDetailsByIDMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMariaDB.GetDetailsByID(ctx, "1")
	}
}

func BenchmarkGetDetailsByIDMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetDetailsByID(ctx, "6477c2ffc1adaf0e33063c09")
	}
}
