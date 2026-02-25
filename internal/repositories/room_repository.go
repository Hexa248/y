package repositories

import (
	"hotel-booking/database"
	"hotel-booking/internal/models"
)

type RoomRepository struct{ DB *database.DB }

func (r RoomRepository) ByHotelID(hotelID int) []models.Room {
	rooms := []models.Room{}
	for _, room := range r.DB.Rooms {
		if room.HotelID == hotelID {
			rooms = append(rooms, room)
		}
	}
	return rooms
}
