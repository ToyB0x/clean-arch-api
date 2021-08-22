package mysql

import (
	"context"
	"errors"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store/mysql/models"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/toaru/clean-arch-api/pkg/server/infra/store"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
	"github.com/toaru/clean-arch-api/pkg/server/domain/repository"
)

func NewReservationRepository(sqlHandler store.SqlHandler) repository.ReservationRepository {
	reservationRepository := ReservationStore{sqlHandler}
	return &reservationRepository
}

type ReservationStore struct {
	store.SqlHandler
}

func (s *ReservationStore) CreatIfAvailable(ctx context.Context, reservation *model.Reservation) error {
	tx, err := s.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	schedule, err := models.FindSchedule(ctx, s.Conn, reservation.ScheduleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// check reservation availability
	if uint(schedule.Stock) == 0 {
		tx.Rollback()
		return errors.New("no stock")
	}

	r := models.Reservation{
		ID:         reservation.ID,
		ScheduleID: schedule.ID,
		TicketID:   reservation.TicketID,
	}

	if err = r.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}

	schedule.Stock -= 1
	if _, err := schedule.Update(ctx, tx, boil.Whitelist(models.ScheduleColumns.Stock, models.ScheduleColumns.UpdatedAt)); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
