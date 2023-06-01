package sql

import (
	"context"
	"mongo-vs-postgres/internal/entities"
	"mongo-vs-postgres/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type postgresImpl struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) repositories.RepositoriesDoer {
	return postgresImpl{conn}
}

func (p postgresImpl) GetUsers(ctx context.Context) ([]entities.User, error) {
	users := make([]entities.User, 0)

	err := p.conn.SelectContext(ctx, &users, `
		SELECT 
			users.id,
			users.name,
			cities.name as city,
			states.name as state,
			states.country as country
		FROM users 
		INNER JOIN cities ON users.city_id = cities.id
		INNER JOIN states ON cities.state_id = states.id
	`)

	return users, err
}

func (p postgresImpl) GetUserByID(ctx context.Context, id string) (entities.User, error) {
	var user entities.User

	err := p.conn.GetContext(ctx, &user, `
		SELECT 
			users.id,
			users.name,
			cities.name as city,
			states.name as state,
			states.country as country
		FROM users 
		INNER JOIN cities ON users.city_id = cities.id
		INNER JOIN states ON cities.state_id = states.id
		WHERE users.id = ?
	`, id)

	if err != nil {
		return user, err
	}

	return user, nil
}
