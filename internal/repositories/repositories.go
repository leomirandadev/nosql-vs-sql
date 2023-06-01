package repositories

import (
	"context"
	"nosql-vs-sql/internal/entities"
)

type RepositoriesDoer interface {
	GetUsers(ctx context.Context) ([]entities.UserDetails, error)
	GetUserByID(ctx context.Context, id string) (entities.UserDetails, error)
}
