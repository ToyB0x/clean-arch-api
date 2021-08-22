package mock

import (
	"context"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

type ReservationStore struct {
	OnCreatIfAvailable func(ctx context.Context, reservation *model.Reservation) error
}

func (s *ReservationStore) CreatIfAvailable(ctx context.Context, reservation *model.Reservation) error {
	return s.OnCreatIfAvailable(ctx, reservation)
}
