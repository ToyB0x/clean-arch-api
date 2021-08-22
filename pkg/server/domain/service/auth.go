package service

// 政府提供の外部認証システムを想定
type AuthService interface {
	IsValidTicket(ticketID string) bool
}
