package exec

import (
	"context"
	"log"
	"nosql-vs-sql/internal/usecases"
	"time"
)

func Run(useCases usecases.UseCases, userID, stateID string) {
	ctx := context.Background()

	log.Println("------ USER ------")

	now := time.Now()
	users, err := useCases.Users.GetDetailsAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), "qty users:", len(users))

	now = time.Now()
	user, err := useCases.Users.GetDetailsByID(ctx, userID)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), user)

	log.Println("------ STATE ------")

	now = time.Now()
	states, err := useCases.States.GetAll(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), "qty states:", len(states))

	now = time.Now()
	state, err := useCases.States.GetByID(ctx, stateID)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		return
	}
	log.Println(time.Since(now), state)
}
