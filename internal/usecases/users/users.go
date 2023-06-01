package users

import (
	"context"
	"errors"
	"nosql-vs-sql/internal/entities"
	"nosql-vs-sql/internal/repositories"
)

type usersUseCase struct {
	repo repositories.RepositoriesDoer
}

type UsersUseCase interface {
	GetDetailsAll(ctx context.Context) (users []entities.UserDetails, err error)
	GetDetailsByID(ctx context.Context, id string) (user entities.UserDetails, err error)
}

func New(repo repositories.RepositoriesDoer) UsersUseCase {
	return usersUseCase{repo}
}

func (u usersUseCase) GetDetailsAll(ctx context.Context) (users []entities.UserDetails, err error) {
	users, err = u.repo.GetUsersDetails(ctx)
	return users, err
}

func (u usersUseCase) GetDetailsByID(ctx context.Context, id string) (user entities.UserDetails, err error) {
	if id == "" {
		return user, errors.New("you must to send the id")
	}

	user, err = u.repo.GetUserDetailsByID(ctx, id)

	return user, err
}
