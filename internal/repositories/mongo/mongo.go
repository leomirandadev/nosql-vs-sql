package mongo

import (
	"context"
	"mongo-vs-postgres/internal/entities"
	"mongo-vs-postgres/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoImpl struct {
	conn *mongo.Database
}

func New(conn *mongo.Database) repositories.RepositoriesDoer {
	return mongoImpl{conn}
}

func (m mongoImpl) GetUsers(ctx context.Context) ([]entities.User, error) {
	result := make([]entities.User, 0)

	// TODO complete pipeline to aggregate collections
	cur, err := m.conn.Collection("users").Aggregate(ctx, mongo.Pipeline{})
	if err != nil {
		return result, err
	}
	defer cur.Close(ctx)

	if err := cur.Err(); err != nil {
		return result, err
	}

	if err = cur.All(ctx, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (m mongoImpl) GetUserByID(ctx context.Context, id string) (user entities.User, err error) {

	getByID := bson.D{{
		Key:   "$match",
		Value: bson.M{"_id": id},
	}}

	// TODO complete pipeline to aggregate collections
	cur, err := m.conn.Collection("users").Aggregate(ctx, mongo.Pipeline{
		getByID,
	})

	if err != nil {
		return user, err
	}
	defer cur.Close(ctx)

	if err := cur.Err(); err != nil {
		return user, err
	}

	if err = cur.All(ctx, &user); err != nil {
		return user, err
	}

	return user, nil
}
