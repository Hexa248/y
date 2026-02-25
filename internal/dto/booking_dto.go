package dto

type BookingRequest struct {
	HotelID  int
	RoomID   int
	CheckIn  string
	CheckOut string
}
