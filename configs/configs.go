package configs

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("connection mongo")
	}

	return client
}

func NewPostgresConn() *sqlx.DB {
	conn, err := sqlx.Connect("postgres", "user=root password=root dbname=mongovspostgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("[ERROR]", err)
	}

	return conn
}
