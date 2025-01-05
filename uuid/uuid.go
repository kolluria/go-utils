package uuid

import (
	"context"

	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func (u UUID) String() string {
	return u.UUID.String()
}

func NewV7(ctx context.Context) UUID {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return UUID{id}
}

func MustParse(s string) UUID {
	return UUID{uuid.MustParse(s)}
}
