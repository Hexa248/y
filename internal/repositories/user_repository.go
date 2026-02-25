package repositories

import (
	"hotel-booking/database"
	"hotel-booking/internal/models"
)

type UserRepository struct{ DB *database.DB }

func (r UserRepository) FindByEmail(email string) (models.User, bool) {
	for _, u := range r.DB.Users {
		if u.Email == email {
			return u, true
		}
	}
	return models.User{}, false
}
