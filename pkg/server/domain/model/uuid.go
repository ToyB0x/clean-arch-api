package model

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() uuid.UUID
}
