package repositories

import (
	"context"
	"nosql-vs-sql/internal/entities"
)

type RepositoriesDoer interface {
	GetStates(ctx context.Context) ([]entities.State, error)
	GetStateByID(ctx context.Context, id string) (entities.State, error)
	GetUsersDetails(ctx context.Context) ([]entities.UserDetails, error)
	GetUserDetailsByID(ctx context.Context, id string) (entities.UserDetails, error)
}
