package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/domain/service"
	mock2 "github.com/toaru/clean-arch-api/pkg/server/infra/memstore/mock"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mock"

	"github.com/toaru/clean-arch-api/pkg/server/usecase"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

func Test_NewScheduleUsecase(t *testing.T) {
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

func Test_scheduleUsecase_GetByMonth(t *testing.T) {
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

func Test_scheduleUsecase_GetByID(t *testing.T) {
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
	type fields struct {
		scheduleRepo    repository.ScheduleRepository
		reservationRepo repository.ReservationRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Schedule
	}{
		{"success",
			fields{
				scheduleRepo: &mock.ScheduleStore{
					OnFindByID: func(ctx context.Context, id string) (*model.Schedule, error) {
						return &schedule, nil
					},
				},
			},
			args{
				ctx: context.Background(),
				id:  "1",
			},
			&schedule,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase.ScheduleUsecaseExport{
				ScheduleRepo:    tt.fields.scheduleRepo,
				ReservationRepo: tt.fields.reservationRepo,
			}
			if got, _ := u.GetByID(tt.args.ctx, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID = %v, want %v", got, tt.want)
			}
		})
	}
}
