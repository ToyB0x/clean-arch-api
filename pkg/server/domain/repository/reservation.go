package repository

import (
	"context"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

type ReservationRepository interface {
	CreatIfAvailable(ctx context.Context, reservation *model.Reservation) error
}
