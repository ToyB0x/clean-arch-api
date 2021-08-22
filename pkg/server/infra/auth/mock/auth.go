package mock

type AuthService struct {
	OnIsValidTicket func(ticketID string) bool
}

func (s *AuthService) IsValidTicket(ticketID string) bool {
	return s.OnIsValidTicket(ticketID)
}
