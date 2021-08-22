package auth_test

import (
	"reflect"
	"testing"

	"github.com/toaru/clean-arch-api/pkg/server/infra/auth"

	"github.com/toaru/clean-arch-api/pkg/server/domain/service"
)

func TestNewAuthService(t *testing.T) {
	tests := []struct {
		name string
		want service.AuthService
	}{
		{
			"success",
			&auth.AuthService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auth.NewAuthService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_IsValidTicket(t *testing.T) {
	type args struct {
		ticketID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"always valid",
			args{
				ticketID: "tokyo-abcde-12345",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auth.AuthService{}
			if got := a.IsValidTicket(tt.args.ticketID); got != tt.want {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}
