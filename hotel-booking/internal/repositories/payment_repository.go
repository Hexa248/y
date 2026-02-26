package repositories

import "hotel-booking/internal/models"

type PaymentRepository struct{ payments []models.Payment }

func NewPaymentRepository() *PaymentRepository { return &PaymentRepository{} }

func (r *PaymentRepository) Create(p models.Payment) models.Payment {
	p.ID = len(r.payments) + 1
	r.payments = append(r.payments, p)
	return p
}
