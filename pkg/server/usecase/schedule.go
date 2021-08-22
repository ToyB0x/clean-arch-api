package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/toaru/clean-arch-api/pkg/server/domain/service"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

type ScheduleUsecase interface {
	GetByMonth(ctx context.Context, year, month int) ([]model.Schedule, error)
	GetByMonthMemStore(ctx context.Context, year, month int) ([]model.Schedule, error)
	GetByID(ctx context.Context, scheduleID string) (*model.Schedule, error)
}

type scheduleUsecase struct {
	ScheduleRepo    repository.ScheduleRepository
	ReservationRepo repository.ReservationRepository
	MemStoreService service.MemStoreService
}

func NewScheduleUsecase(scheduleRepo repository.ScheduleRepository, reservationRepo repository.ReservationRepository, memStoreService service.MemStoreService) ScheduleUsecase {
	scheduleUsecase := scheduleUsecase{scheduleRepo, reservationRepo, memStoreService}
	return &scheduleUsecase
}

func (u *scheduleUsecase) GetByMonth(ctx context.Context, year, month int) ([]model.Schedule, error) {
	return u.ScheduleRepo.FindByMonth(ctx, year, month)
}

func (u *scheduleUsecase) GetByMonthMemStore(ctx context.Context, year, month int) ([]model.Schedule, error) {
	key := fmt.Sprintf("%d/%d", year, month)
	jsonCached, err := u.MemStoreService.Get(key)

	// get from db
	if err != nil {
		schedules, err := u.GetByMonth(ctx, year, month)
		if err != nil {
			return nil, err
		}
		// set new cache
		if schedules != nil {
			j, err := json.Marshal(schedules)
			if err != nil {
				return nil, err
			}
			err = u.MemStoreService.Add(key, j, 1)
		}
		return schedules, err
	}

	// use existence cache
	var schedules []model.Schedule
	err = json.Unmarshal(jsonCached, &schedules)
	return schedules, err
}

func (u *scheduleUsecase) GetByID(ctx context.Context, scheduleID string) (*model.Schedule, error) {
	return u.ScheduleRepo.FindByID(ctx, scheduleID)
}
