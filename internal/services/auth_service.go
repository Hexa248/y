package services

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"hotel-booking/pkg/security"
)

type AuthService struct{ Users repositories.UserRepository }

func (s AuthService) Login(email, password string) (models.User, string, bool) {
	u, ok := s.Users.FindByEmail(email)
	if !ok || !security.CheckPassword(u.Password, password) {
		return models.User{}, "", false
	}
	return u, security.GenerateToken(email), true
}
