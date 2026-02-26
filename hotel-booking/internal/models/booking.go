package models

import "time"

type BookingStatus string

const (
	BookingPending   BookingStatus = "pending"
	BookingConfirmed BookingStatus = "confirmed"
)

type Booking struct {
	ID         int
	UserID     int
	RoomID     int
	CheckIn    time.Time
	CheckOut   time.Time
	GuestCount int
	TotalPrice int
	Status     BookingStatus
}
