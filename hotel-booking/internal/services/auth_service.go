package services

import (
	"errors"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type AuthService struct{ users *repositories.UserRepository }

func NewAuthService(users *repositories.UserRepository) *AuthService {
	return &AuthService{users: users}
}

func (s *AuthService) Login(email, password string) (models.User, error) {
	u, ok := s.users.FindByEmail(email)
	if !ok || u.Password != password {
		return models.User{}, errors.New("invalid credentials")
	}
	return u, nil
}

func (s *AuthService) Register(name, email, password, role string) models.User {
	return s.users.Create(models.User{Name: name, Email: email, Password: password, Role: models.Role(role)})
}
