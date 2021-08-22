package model_test

import (
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/domain/model"
)

type mockAuthService struct{}

func (mockAuthService) IsValidTicket(ticketID string) bool {
	if ticketID == "valid" {
		return true
	} else {
		return false
	}
}

func TestNewReservation(t *testing.T) {
	type args struct {
		ticketID   string
		scheduleID string
	}
	tests := []struct {
		name string
		args args
		want *model.Reservation
	}{
		{
			name: "with valid ticket",
			args: args{
				ticketID:   "valid",
				scheduleID: "1",
			},
			want: &model.Reservation{
				ScheduleID: "1",
				ID:         "00000000-0000-0000-0000-000000000000",
				TicketID:   "valid",
			},
		},
		{
			name: "with blank ticket id",
			args: args{
				ticketID:   "not valid",
				scheduleID: "1",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := model.NewReservation(tt.args.ticketID, tt.args.scheduleID, mockAuthService{}, mockUUIDGen{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
