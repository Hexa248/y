package services

import (
	"fmt"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
)

type PaymentService struct {
	repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Pay(bookingID int, method string) models.Payment {
	return s.repo.Create(models.Payment{BookingID: bookingID, Method: models.PaymentMethod(method), Status: "paid", Reference: fmt.Sprintf("PAY-%d", bookingID)})
}
