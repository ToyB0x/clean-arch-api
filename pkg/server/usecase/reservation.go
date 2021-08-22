package usecase

import (
	"context"

	"github.com/toaru/clean-arch-api/pkg/server/domain/service"

	"github.com/google/uuid"
	"github.com/toaru/clean-arch-api/pkg/server/infra/auth"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"

	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

type ReservationUsecase interface {
	Create(ctx context.Context, ticketID, scheduleID string) error
}

type reservationUsecase struct {
	ReservationRepo repository.ReservationRepository
	AuthService     service.AuthService
}

func NewReservationUsecase(reservationRepo repository.ReservationRepository, authService service.AuthService) ReservationUsecase {
	reservationUsecase := reservationUsecase{reservationRepo, authService}
	return &reservationUsecase
}

type uuidGen struct{}

func (uuidGen) New() uuid.UUID {
	return uuid.New()
}

func (u *reservationUsecase) Create(ctx context.Context, ticketID, scheduleID string) error {
	authService := auth.NewAuthService()
	reservation, err := model.NewReservation(ticketID, scheduleID, authService, uuidGen{})
	if err != nil {
		return err
	}
	return u.ReservationRepo.CreatIfAvailable(ctx, reservation)
}
