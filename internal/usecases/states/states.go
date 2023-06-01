package states

import (
	"context"
	"errors"
	"nosql-vs-sql/internal/entities"
	"nosql-vs-sql/internal/repositories"
)

type statesUseCase struct {
	repo repositories.RepositoriesDoer
}

type StatesUseCase interface {
	GetAll(ctx context.Context) (states []entities.State, err error)
	GetByID(ctx context.Context, id string) (state entities.State, err error)
}

func New(repo repositories.RepositoriesDoer) StatesUseCase {
	return statesUseCase{repo}
}

func (u statesUseCase) GetAll(ctx context.Context) (states []entities.State, err error) {
	states, err = u.repo.GetStates(ctx)
	return states, err
}

func (u statesUseCase) GetByID(ctx context.Context, id string) (state entities.State, err error) {
	if id == "" {
		return state, errors.New("you must to send the id")
	}

	state, err = u.repo.GetStateByID(ctx, id)

	return state, err
}
