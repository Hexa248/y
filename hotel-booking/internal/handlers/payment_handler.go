package handlers

import (
	"net/http"
	"strconv"

	"hotel-booking/internal/services"
)

type PaymentHandler struct {
	renderer *Renderer
	payment  *services.PaymentService
}

func NewPaymentHandler(r *Renderer, p *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{renderer: r, payment: p}
}

func (h *PaymentHandler) Page(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("booking_id"))
	h.renderer.Render(w, "payment/payment.html", map[string]any{"Title": "Pembayaran", "BookingID": id})
}

func (h *PaymentHandler) Pay(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("booking_id"))
	h.payment.Pay(id, r.FormValue("method"))
	http.Redirect(w, r, "/booking/history", http.StatusSeeOther)
}
