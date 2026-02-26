package dto

type BookingRequest struct {
	RoomID     int    `json:"room_id"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
	GuestCount int    `json:"guest_count"`
}
