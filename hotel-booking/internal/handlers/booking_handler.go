package handlers

import (
	"net/http"
	"strconv"

	"hotel-booking/internal/services"
)

type BookingHandler struct {
	renderer *Renderer
	booking  *services.BookingService
}

func NewBookingHandler(r *Renderer, b *services.BookingService) *BookingHandler {
	return &BookingHandler{renderer: r, booking: b}
}

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	roomID, _ := strconv.Atoi(r.FormValue("room_id"))
	guest, _ := strconv.Atoi(r.FormValue("guest_count"))
	total, _ := strconv.Atoi(r.FormValue("total_price"))
	b := h.booking.Create(1, roomID, r.FormValue("check_in"), r.FormValue("check_out"), guest, total)
	http.Redirect(w, r, "/payment?booking_id="+strconv.Itoa(b.ID), http.StatusSeeOther)
}

func (h *BookingHandler) History(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "booking/history.html", map[string]any{"Title": "Riwayat Booking", "Bookings": h.booking.ByUser(1)})
}
