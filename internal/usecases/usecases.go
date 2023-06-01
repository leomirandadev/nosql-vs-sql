package usecases

import (
	"nosql-vs-sql/internal/repositories"
	"nosql-vs-sql/internal/usecases/states"
	"nosql-vs-sql/internal/usecases/users"
)

type UseCases struct {
	Users  users.UsersUseCase
	States states.StatesUseCase
}

func New(repo repositories.RepositoriesDoer) UseCases {
	return UseCases{
		Users:  users.New(repo),
		States: states.New(repo),
	}
}
