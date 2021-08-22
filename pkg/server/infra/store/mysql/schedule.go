package mysql

import (
	"context"
	"time"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql/models"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

func NewScheduleRepository(sqlHandler store.SqlHandler) repository.ScheduleRepository {
	scheduleRepository := ScheduleStore{sqlHandler}
	return &scheduleRepository
}

type ScheduleStore struct {
	store.SqlHandler
}

func (s *ScheduleStore) Save(ctx context.Context, schedule *model.Schedule) error {
	d := time.Date(int(schedule.Year), time.Month(schedule.Month), int(schedule.Day), 0, 0, 0, 0, time.UTC)
	f := models.Schedule{
		ID:           schedule.ID,
		Date:         d,
		Hour:         uint8(schedule.Hour),
		Min:          uint8(schedule.Min),
		MaxAvailable: schedule.MaxAvailable,
		Stock:        schedule.Stock,
	}
	err := f.Insert(ctx, s.Conn, boil.Infer())
	return err
}

func (s *ScheduleStore) FindByMonth(ctx context.Context, year, month int) ([]model.Schedule, error) {
	d1 := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	results, err := models.Schedules(
		models.ScheduleWhere.Date.GTE(d1),
		models.ScheduleWhere.Date.LT(d2),
	).All(ctx, s.Conn)

	if err != nil {
		return nil, err
	}

	schedules := make([]model.Schedule, 0, len(results))
	for _, r := range results {
		schedule := model.Schedule{
			ID:           r.ID,
			MaxAvailable: r.MaxAvailable,
			Stock:        r.Stock,
			ScheduleDate: model.ScheduleDate{
				Year:  uint(r.Date.Year()),
				Month: uint(r.Date.Month()),
				Day:   uint(r.Date.Day()),
				Hour:  uint(r.Hour),
				Min:   uint(r.Min),
			},
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *ScheduleStore) FindByID(ctx context.Context, scheduleID string) (*model.Schedule, error) {
	r, err := models.FindSchedule(ctx, s.Conn, scheduleID)
	if err != nil {
		return nil, err
	}

	return &model.Schedule{
		ID:           r.ID,
		Stock:        r.Stock,
		MaxAvailable: r.MaxAvailable,
		ScheduleDate: model.ScheduleDate{
			Year:  uint(r.Date.Year()),
			Month: uint(r.Date.Month()),
			Day:   uint(r.Date.Day()),
			Hour:  uint(r.Hour),
			Min:   uint(r.Min),
		},
	}, nil
}
