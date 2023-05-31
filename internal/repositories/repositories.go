package repositories

import (
	"context"
	"mongo-vs-postgres/internal/entities"
)

type RepositoriesDoer interface {
	GetUsers(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context, id string) (entities.User, error)
}
