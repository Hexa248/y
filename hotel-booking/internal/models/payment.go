package models

type PaymentMethod string

const (
	PaymentVA     PaymentMethod = "virtual_account"
	PaymentQRIS   PaymentMethod = "qris"
	PaymentCard   PaymentMethod = "credit_card"
	PaymentWallet PaymentMethod = "e_wallet"
)

type Payment struct {
	ID        int
	BookingID int
	Method    PaymentMethod
	Status    string
	Reference string
}
