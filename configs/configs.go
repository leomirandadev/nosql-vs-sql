package configs

import (
	"context"
	"database/sql"
	"log"
	"time"

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

func NewPostgresConn() *sql.DB {
	conn, err := sql.Open("postgres", "user=root password=root dbname=mongovspostgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("[ERROR]", err)
	}
	return conn
}
