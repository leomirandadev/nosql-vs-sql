package users

import (
	"context"
	"mongo-vs-postgres/configs"
	mongoRepo "mongo-vs-postgres/internal/repositories/mongo"
	"mongo-vs-postgres/internal/repositories/postgres"
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
var repoPostgres = postgres.New(postgressConn)
var usersUseCasePostgres = New(repoPostgres)

func BenchmarkGetAllPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetAll(ctx)
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

func BenchmarkGetByIDMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetByID(ctx, "6477c2ffc1adaf0e33063c09")
	}
}
