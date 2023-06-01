package mongo

import (
	"context"
	"errors"
	"nosql-vs-sql/internal/entities"
	"nosql-vs-sql/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoImpl struct {
	conn *mongo.Database
}

func New(conn *mongo.Database) repositories.RepositoriesDoer {
	return mongoImpl{conn}
}

func (m mongoImpl) GetUsersDetails(ctx context.Context) ([]entities.UserDetails, error) {
	result := make([]entities.UserDetails, 0)

	cur, err := m.conn.Collection("users").Aggregate(ctx, mongo.Pipeline{
		lookupCity,
		unwindCity,
		lookupState,
		unwindState,
		group,
		unwindCity,
		unwindState,
		unwindName,
		unwindCountry,
	})
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

func (m mongoImpl) GetUserDetailsByID(ctx context.Context, id string) (user entities.UserDetails, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, errors.New("invalid id")
	}

	getByID := bson.D{{
		Key:   "$match",
		Value: bson.M{"_id": objectId},
	}}

	cur, err := m.conn.Collection("users").Aggregate(ctx, mongo.Pipeline{
		getByID,
		lookupCity,
		unwindCity,
		lookupState,
		unwindState,
		group,
		unwindCity,
		unwindState,
		unwindName,
		unwindCountry,
	})

	if err != nil {
		return user, err
	}
	defer cur.Close(ctx)

	if err := cur.Err(); err != nil {
		return user, err
	}

	var results []entities.UserDetails
	if err = cur.All(ctx, &results); err != nil {
		return user, err
	}

	if len(results) < 1 {
		return user, errors.New("user not found")
	}

	return results[0], nil
}

func (m mongoImpl) GetStates(ctx context.Context) ([]entities.State, error) {
	result := make([]entities.State, 0)

	cur, err := m.conn.Collection("states").Find(ctx, bson.M{})
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

func (m mongoImpl) GetStateByID(ctx context.Context, id string) (user entities.State, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, errors.New("invalid id")
	}

	cur := m.conn.Collection("states").FindOne(ctx, bson.M{"_id": objectId}, nil)

	if err := cur.Err(); err != nil {
		return user, err
	}

	var result entities.State
	if err = cur.Decode(&result); err != nil {
		return user, err
	}

	return result, nil
}

var lookupCity = bson.D{{
	Key: "$lookup",
	Value: bson.M{
		"from":         "cities",
		"localField":   "city_id",
		"foreignField": "_id",
		"as":           "city",
	},
}}

var unwindCity = bson.D{{
	Key:   "$unwind",
	Value: "$city",
}}

var lookupState = bson.D{{
	Key: "$lookup",
	Value: bson.M{
		"from":         "states",
		"localField":   "city.state_id",
		"foreignField": "_id",
		"as":           "state",
	},
}}

var unwindState = bson.D{{
	Key:   "$unwind",
	Value: "$state",
}}

var group = bson.D{{
	Key: "$group",
	Value: bson.M{
		"_id": "$_id",
		"name": bson.M{
			"$addToSet": "$name",
		},
		"city": bson.M{
			"$addToSet": "$city.name",
		},
		"state": bson.M{
			"$addToSet": "$state.name",
		},
		"country": bson.M{
			"$addToSet": "$state.country",
		},
	},
}}

var unwindName = bson.D{{
	Key:   "$unwind",
	Value: "$name",
}}

var unwindCountry = bson.D{{
	Key:   "$unwind",
	Value: "$country",
}}

/*

db.users.aggregate([
	{
		$lookup: {
			from: "cities",
			localField:"city_id",
			foreignField: "_id",
			as: "city",
		},
	},
	{
		$unwind:"$city"
	},
	{
		$lookup: {
			from: "states",
			localField: "city.state_id",
			foreignField: "_id",
			as: "state",
		},
	},
	{
		$unwind:"$state"
	},
	{
		$group: {
			_id: "$_id",
			name: {
				$addToSet: "$name"
			},
			city: {
				$addToSet: "$city.name"
			},
			state: {
				$addToSet: "$state.name"
			},
			country: {
				$addToSet: "$state.country"
			},
		}
	},
	{
		$unwind:"$state",
	},
	{
		$unwind:"$city",
	},
	{
		$unwind:"$name",
	},
	{
		$unwind:"$country",
	},
])
*/
