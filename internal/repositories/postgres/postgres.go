package postgres

import (
	"context"
	"database/sql"
	"log"
	"mongo-vs-postgres/internal/entities"
	"mongo-vs-postgres/internal/repositories"
)

type postgresImpl struct {
	conn *sql.DB
}

func New(conn *sql.DB) repositories.RepositoriesDoer {
	return postgresImpl{conn}
}

func (p postgresImpl) GetUsers(ctx context.Context) ([]entities.User, error) {
	users := make([]entities.User, 0)

	rows, err := p.conn.QueryContext(ctx, `
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

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user entities.User

		err := rows.Scan(&user.ID, &user.Name, &user.City, &user.State, &user.Country)
		if err != nil {
			log.Printf("error %v", err)
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (p postgresImpl) GetUserByID(ctx context.Context, id string) (entities.User, error) {
	var user entities.User

	query := p.conn.QueryRowContext(ctx, `
		SELECT 
			users.id,
			users.name,
			cities.name as city,
			states.name as state,
			states.country as country
		FROM users 
		INNER JOIN cities ON users.city_id = cities.id
		INNER JOIN states ON cities.state_id = states.id
		WHERE users.id = $1
	`, id)

	if err := query.Err(); err != nil {
		log.Println("here")
		return user, err
	}

	err := query.Scan(&user.ID, &user.Name, &user.City, &user.State, &user.Country)
	if err != nil {
		return user, err
	}

	return user, nil
}
