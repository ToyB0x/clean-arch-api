package auth

import (
	"github.com/toaru/clean-arch-api/pkg/server/domain/service"
)

type authService struct{}

func NewAuthService() service.AuthService {
	authService := &authService{}
	return authService
}

func (a *authService) IsValidTicket(ticketID string) bool {
	// the governmental auth service mock
	return true
}
