package sql

import (
	"context"
	"nosql-vs-sql/internal/entities"
	"nosql-vs-sql/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type sqlImpl struct {
	conn       *sqlx.DB
	isPostgres bool
}

func New(conn *sqlx.DB, isPostgres bool) repositories.RepositoriesDoer {
	return sqlImpl{conn, isPostgres}
}

func (p sqlImpl) GetUsers(ctx context.Context) ([]entities.UserDetails, error) {
	users := make([]entities.UserDetails, 0)

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

func (p sqlImpl) GetUserByID(ctx context.Context, id string) (entities.UserDetails, error) {
	var user entities.UserDetails

	bindVar := "?"
	if p.isPostgres {
		bindVar = "$1"
	}

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
		WHERE users.id = `+bindVar+`
	`, id)

	if err != nil {
		return user, err
	}

	return user, nil
}
