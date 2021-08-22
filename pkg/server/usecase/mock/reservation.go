package mock

import (
	"context"
)

type ReservationUsecase struct {
	OnCreate func(ctx context.Context, ticketID, scheduleID string) error
}

func (u *ReservationUsecase) Create(ctx context.Context, ticketID, scheduleID string) error {
	return u.OnCreate(ctx, ticketID, scheduleID)
}
