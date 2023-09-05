package users

import (
	"context"
	"nosql-vs-sql/configs"
	mongoRepo "nosql-vs-sql/internal/repositories/mongo"
	"nosql-vs-sql/internal/repositories/sql"
	"testing"
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

func BenchmarkUsersGetDetailsAllPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetDetailsAll(ctx)
	}
}

func BenchmarkUsersGetDetailsAllMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMariaDB.GetDetailsAll(ctx)
	}
}

func BenchmarkUsersGetDetailsAllMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetDetailsAll(ctx)
	}
}

func BenchmarkUsersGetDetailsByIDPostgres(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCasePostgres.GetDetailsByID(ctx, "1")
	}
}

func BenchmarkUsersGetDetailsByIDMariaDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMariaDB.GetDetailsByID(ctx, "1")
	}
}

func BenchmarkUsersGetDetailsByIDMongo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usersUseCaseMongo.GetDetailsByID(ctx, "6477c2ffc1adaf0e33063c09")
	}
}
