package repositories

import (
	"hotel-booking/database"
	"hotel-booking/internal/models"
)

type HotelRepository struct{ DB *database.DB }

func (r HotelRepository) All() []models.Hotel { return r.DB.Hotels }
func (r HotelRepository) ByID(id int) (models.Hotel, bool) {
	for _, h := range r.DB.Hotels {
		if h.ID == id {
			return h, true
		}
	}
	return models.Hotel{}, false
}
