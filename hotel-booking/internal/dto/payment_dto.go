package dto

type PaymentRequest struct {
	BookingID int    `json:"booking_id"`
	Method    string `json:"method"`
}
