package services

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type HotelService struct {
	Hotels repositories.HotelRepository
	Rooms  repositories.RoomRepository
}

func (s HotelService) HotelsWithRooms() []models.Hotel { return s.Hotels.All() }
func (s HotelService) HotelDetail(id int) (models.Hotel, []models.Room, bool) {
	h, ok := s.Hotels.ByID(id)
	if !ok {
		return models.Hotel{}, nil, false
	}
	return h, s.Rooms.ByHotelID(id), true
}
