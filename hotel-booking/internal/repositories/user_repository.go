package repositories

import "hotel-booking/internal/models"

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: []models.User{{ID: 1, Name: "Admin", Email: "admin@nusa.id", Password: "admin123", Role: models.RoleAdmin}, {ID: 2, Name: "Staff", Email: "staff@nusa.id", Password: "staff123", Role: models.RoleStaff}}}
}

func (r *UserRepository) FindByEmail(email string) (models.User, bool) {
	for _, u := range r.users {
		if u.Email == email {
			return u, true
		}
	}
	return models.User{}, false
}

func (r *UserRepository) Create(u models.User) models.User {
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return u
}
