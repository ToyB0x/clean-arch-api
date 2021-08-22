package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/usecase"

	"github.com/toaru/clean-arch-api/pkg/server/domain/service"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"

	amock "github.com/toaru/clean-arch-api/pkg/server/infra/auth/mock"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mock"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

func TestNewReservationUsecase(t *testing.T) {
	type args struct {
		reservationRepo repository.ReservationRepository
		authService     service.AuthService
	}
	tests := []struct {
		name string
		args args
		want usecase.ReservationUsecase
	}{
		{
			name: "success",
			args: args{&mock.ReservationStore{}, &amock.AuthService{}},
			want: &usecase.ReservationUsecaseExport{ReservationRepo: &mock.ReservationStore{}, AuthService: &amock.AuthService{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := usecase.NewReservationUsecase(tt.args.reservationRepo, tt.args.authService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReservationUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reservationUsecase_Create(t *testing.T) {
	type fields struct {
		Repo repository.ReservationRepository
	}
	type args struct {
		ctx                  context.Context
		ticketID, scheduleID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"success",
			fields{Repo: &mock.ReservationStore{
				OnCreatIfAvailable: func(ctx context.Context, reservation *model.Reservation) error {
					return nil
				},
			}},
			args{
				ctx:        context.Background(),
				ticketID:   "ticket-1",
				scheduleID: "1",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase.ReservationUsecaseExport{
				ReservationRepo: tt.fields.Repo,
			}
			if err := u.Create(tt.args.ctx, tt.args.ticketID, tt.args.scheduleID); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
