package users

import (
	"context"
	"errors"
	"mongo-vs-postgres/internal/entities"
	"mongo-vs-postgres/internal/repositories"
)

type usersUseCase struct {
	repo repositories.RepositoriesDoer
}

func New(repo repositories.RepositoriesDoer) usersUseCase {
	return usersUseCase{repo}
}

func (u usersUseCase) GetAll(ctx context.Context) (users []entities.User, err error) {
	users, err = u.repo.GetUsers(ctx)
	return users, err
}

func (u usersUseCase) GetByID(ctx context.Context, id string) (user entities.User, err error) {
	if id == "" {
		return user, errors.New("you must to send the id")
	}

	user, err = u.repo.GetUserByID(ctx, id)

	return user, err
}
