package mock

import (
	"context"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

type ScheduleUsecase struct {
	OnGetByMonth         func(ctx context.Context, year, month int) ([]model.Schedule, error)
	OnGetByMonthMemStore func(ctx context.Context, year, month int) ([]model.Schedule, error)
	OnGetByID            func(ctx context.Context, scheduleID string) (*model.Schedule, error)
}

func (u *ScheduleUsecase) GetByMonth(ctx context.Context, year, month int) ([]model.Schedule, error) {
	return u.OnGetByMonth(ctx, year, month)
}

func (u *ScheduleUsecase) GetByMonthMemStore(ctx context.Context, year, month int) ([]model.Schedule, error) {
	return u.OnGetByMonthMemStore(ctx, year, month)
}

func (u *ScheduleUsecase) GetByID(ctx context.Context, scheduleID string) (*model.Schedule, error) {
	return u.OnGetByID(ctx, scheduleID)
}
