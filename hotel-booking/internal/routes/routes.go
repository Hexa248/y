package routes

import (
	"net/http"

	"hotel-booking/internal/handlers"
	"hotel-booking/internal/middleware"
)

func Register(mux *http.ServeMux, auth *handlers.AuthHandler, hotel *handlers.HotelHandler, room *handlers.RoomHandler, booking *handlers.BookingHandler, payment *handlers.PaymentHandler, admin *handlers.AdminHandler) {
	mux.HandleFunc("/", hotel.Home)
	mux.HandleFunc("/hotels", hotel.CityHotels)
	mux.HandleFunc("/hotel", hotel.Detail)
	mux.HandleFunc("/room", room.Detail)

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			auth.LoginPage(w, r)
			return
		}
		auth.DoLogin(w, r)
	})
	mux.HandleFunc("/register", auth.RegisterPage)
	mux.HandleFunc("/auth/google", auth.GoogleAuth)

	mux.Handle("/booking", middleware.Auth(http.HandlerFunc(booking.Create)))
	mux.Handle("/booking/history", middleware.Auth(http.HandlerFunc(booking.History)))
	mux.Handle("/payment", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			payment.Page(w, r)
			return
		}
		payment.Pay(w, r)
	})))
	mux.Handle("/admin", middleware.Admin(http.HandlerFunc(admin.Dashboard)))
}
