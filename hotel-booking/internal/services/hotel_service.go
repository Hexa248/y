package services

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type HotelService struct{ repo *repositories.HotelRepository }

func NewHotelService(repo *repositories.HotelRepository) *HotelService {
	return &HotelService{repo: repo}
}

func (s *HotelService) Cities() []models.City                  { return s.repo.Cities() }
func (s *HotelService) Hotels() []models.Hotel                 { return s.repo.Hotels() }
func (s *HotelService) HotelsByCity(cityID int) []models.Hotel { return s.repo.HotelsByCity(cityID) }
func (s *HotelService) HotelByID(id int) (models.Hotel, bool)  { return s.repo.HotelByID(id) }
func (s *HotelService) RoomsByHotel(id int) []models.Room      { return s.repo.RoomsByHotel(id) }
func (s *HotelService) RoomByID(id int) (models.Room, bool)    { return s.repo.RoomByID(id) }
