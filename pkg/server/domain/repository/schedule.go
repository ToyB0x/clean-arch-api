package repository

import (
	"context"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

type ScheduleRepository interface {
	Save(ctx context.Context, schedule *model.Schedule) error
	FindByMonth(ctx context.Context, year, month int) ([]model.Schedule, error)
	FindByID(ctx context.Context, scheduleID string) (*model.Schedule, error)
}
