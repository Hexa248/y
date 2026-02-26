package services

import (
	"time"

	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type BookingService struct {
	repo *repositories.BookingRepository
}

func NewBookingService(repo *repositories.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) Create(userID, roomID int, checkIn, checkOut string, guestCount, total int) models.Booking {
	in, _ := time.Parse("2006-01-02", checkIn)
	out, _ := time.Parse("2006-01-02", checkOut)
	return s.repo.Create(models.Booking{UserID: userID, RoomID: roomID, CheckIn: in, CheckOut: out, GuestCount: guestCount, TotalPrice: total, Status: models.BookingConfirmed})
}

func (s *BookingService) ByUser(userID int) []models.Booking { return s.repo.ByUser(userID) }
