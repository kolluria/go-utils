package uuid

import (
	"context"

	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func NewV7(ctx context.Context) (UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return UUID{id}, nil
}
