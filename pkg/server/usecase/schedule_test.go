package usecase_test

import (
	"context"
	"encoding/json"
	"github.com/friendsofgo/errors"
	"reflect"
	"testing"

	mock2 "github.com/toaru/clean-arch-api/pkg/server/infra/memstore/mock"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mock"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
	"github.com/toaru/clean-arch-api/pkg/server/usecase"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
	"github.com/toaru/clean-arch-api/pkg/server/domain/service"
)

func TestNewScheduleUsecase(t *testing.T) {
	type args struct {
		scheduleRepo      repository.ScheduleRepository
		reservationRepo   repository.ReservationRepository
		memoStoreServivce service.MemStoreService
	}
	tests := []struct {
		name string
		args args
		want usecase.ScheduleUsecase
	}{
		{
			name: "success",
			args: args{&mock.ScheduleStore{}, &mock.ReservationStore{}, &mock2.MemStoreService{}},
			want: &usecase.ScheduleUsecaseExport{ScheduleRepo: &mock.ScheduleStore{}, ReservationRepo: &mock.ReservationStore{}, MemStoreService: &mock2.MemStoreService{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := usecase.NewScheduleUsecase(tt.args.scheduleRepo, tt.args.reservationRepo, tt.args.memoStoreServivce); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduleUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetByMonth(t *testing.T) {
	// test data
	schedule := model.Schedule{
		ID:           "1",
		MaxAvailable: 10,
		Stock:        1,
		ScheduleDate: model.ScheduleDate{
			Year:  2020,
			Month: 1,
			Day:   1,
			Hour:  9,
			Min:   20,
		},
	}
	schedules := []model.Schedule{schedule}
	type fields struct {
		scheduleRepo    repository.ScheduleRepository
		reservationRepo repository.ReservationRepository
	}
	type args struct {
		ctx         context.Context
		year, month int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []model.Schedule
	}{
		{"success",
			fields{
				scheduleRepo: &mock.ScheduleStore{
					OnFindByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
						return schedules, nil
					},
				},
			},
			args{
				ctx:   context.Background(),
				year:  2020,
				month: 1,
			},
			schedules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase.ScheduleUsecaseExport{
				ScheduleRepo:    tt.fields.scheduleRepo,
				ReservationRepo: tt.fields.reservationRepo,
			}
			if got, _ := u.GetByMonth(tt.args.ctx, tt.args.year, tt.args.month); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetByMonthMemStore(t *testing.T) {
	// test data
	schedule := model.Schedule{
		ID:           "1",
		MaxAvailable: 10,
		Stock:        1,
		ScheduleDate: model.ScheduleDate{
			Year:  2020,
			Month: 1,
			Day:   1,
			Hour:  9,
			Min:   20,
		},
	}
	schedules := []model.Schedule{schedule}
	schedulesJson, err := json.Marshal(schedules)
	if err != nil {
		panic(err)
	}

	type fields struct {
		scheduleRepo    repository.ScheduleRepository
		reservationRepo repository.ReservationRepository
		memStoreService service.MemStoreService
	}
	type args struct {
		ctx         context.Context
		year, month int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []model.Schedule
	}{
		{"hit cache",
			fields{
				memStoreService: &mock2.MemStoreService{
					OnGet: func(key string) ([]byte, error) {
						return schedulesJson, nil
					},
					OnAdd: func(key string, value []byte, sec int) error {
						return nil
					},
				},
				scheduleRepo: &mock.ScheduleStore{
					OnFindByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
						return schedules, nil
					},
				},
			},
			args{
				ctx:   context.Background(),
				year:  2020,
				month: 1,
			},
			schedules,
		},
		{"no hit cache",
			fields{
				memStoreService: &mock2.MemStoreService{
					OnGet: func(key string) ([]byte, error) {
						return nil, errors.New("no cache")
					},
					OnAdd: func(key string, value []byte, sec int) error {
						return nil
					},
				},
				scheduleRepo: &mock.ScheduleStore{
					OnFindByMonth: func(ctx context.Context, year, month int) ([]model.Schedule, error) {
						return schedules, nil
					},
				},
			},
			args{
				ctx:   context.Background(),
				year:  2020,
				month: 1,
			},
			schedules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase.ScheduleUsecaseExport{
				ScheduleRepo:    tt.fields.scheduleRepo,
				ReservationRepo: tt.fields.reservationRepo,
				MemStoreService: tt.fields.memStoreService,
			}
			if got, _ := u.GetByMonthMemStore(tt.args.ctx, tt.args.year, tt.args.month); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
