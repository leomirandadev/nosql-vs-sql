package users

import (
	"context"
	"mongo-vs-postgres/configs"
	mongoRepo "mongo-vs-postgres/internal/repositories/mongo"
	"mongo-vs-postgres/internal/repositories/sql"
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
var repoPostgres = sql.New(postgressConn)
var usersUseCasePostgres = New(repoPostgres)

// mariadb
var mariadbConn = configs.NewPostgresConn()
var repoMariaDB = sql.New(mariadbConn)
var usersUseCaseMariaDB = New(repoMariaDB)

func BenchmarkGetAllPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetAll(ctx)
	}
}

func BenchmarkGetAllMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMariaDB.GetAll(ctx)
	}
}

func BenchmarkGetAllMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetAll(ctx)
	}
}

func BenchmarkGetByIDPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetByID(ctx, "1")
	}
}

func BenchmarkGetByIDMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMariaDB.GetByID(ctx, "1")
	}
}

func BenchmarkGetByIDMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetByID(ctx, "6477c2ffc1adaf0e33063c09")
	}
}
