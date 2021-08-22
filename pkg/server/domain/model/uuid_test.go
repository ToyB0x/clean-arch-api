package model_test

import "github.com/google/uuid"

type mockUUIDGen struct{}

func (mockUUIDGen) New() uuid.UUID {
	u, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	return u
}
