package repositories

import "hotel-booking/internal/models"

type BookingRepository struct{ bookings []models.Booking }

func NewBookingRepository() *BookingRepository { return &BookingRepository{} }

func (r *BookingRepository) Create(b models.Booking) models.Booking {
	b.ID = len(r.bookings) + 1
	r.bookings = append(r.bookings, b)
	return b
}

func (r *BookingRepository) ByUser(userID int) []models.Booking {
	out := []models.Booking{}
	for _, b := range r.bookings {
		if b.UserID == userID {
			out = append(out, b)
		}
	}
	return out
}
