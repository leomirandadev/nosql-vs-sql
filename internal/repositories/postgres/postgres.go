package postgres

import (
	"context"
	"database/sql"
	"mongo-vs-postgres/internal/entities"
	"mongo-vs-postgres/internal/repositories"
)

type postgresImpl struct {
	conn *sql.DB
}

func New(conn *sql.DB) repositories.RepositoriesDoer {
	return postgresImpl{conn}
}

func (postgresImpl) GetUsers(ctx context.Context) ([]entities.User, error) {
	return []entities.User{}, nil
}

func (postgresImpl) GetUserByID(ctx context.Context, id string) (entities.User, error) {
	return entities.User{}, nil
}
