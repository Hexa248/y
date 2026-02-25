package models

type Payment struct {
	ID            int
	BookingID     int
	Method        string
	Gateway       string
	Status        string
	TransactionID string
}
