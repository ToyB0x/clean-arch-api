package mock

import (
	"context"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

type ScheduleStore struct {
	OnSave        func(ctx context.Context, schedule *model.Schedule) error
	OnFindByMonth func(ctx context.Context, year, month int) ([]model.Schedule, error)
	OnFindByID    func(ctx context.Context, id string) (*model.Schedule, error)
}

func (s *ScheduleStore) Save(ctx context.Context, schedule *model.Schedule) error {
	return s.OnSave(ctx, schedule)
}

func (s *ScheduleStore) FindByMonth(ctx context.Context, year, month int) ([]model.Schedule, error) {
	return s.OnFindByMonth(ctx, year, month)
}
func (s *ScheduleStore) FindByID(ctx context.Context, id string) (*model.Schedule, error) {
	return s.OnFindByID(ctx, id)
}
