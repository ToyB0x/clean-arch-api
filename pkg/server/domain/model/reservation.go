package model

import (
	"github.com/friendsofgo/errors"
	"github.com/toaru/clean-arch-api/pkg/server/domain/service"
)

type Reservation struct {
	ID         string // uuid
	ScheduleID string // uuid
	TicketID   string
}

// ユーザによる予約作成
func NewReservation(ticketID string, scheduleID string, authService service.AuthService, uuidGen UUIDGenerator) (*Reservation, error) {
	if !authService.IsValidTicket(ticketID) {
		return nil, errors.New("invalid ticketID")
	}
	return &Reservation{
		ID:         uuidGen.New().String(),
		ScheduleID: scheduleID,
		TicketID:   ticketID,
	}, nil
}
