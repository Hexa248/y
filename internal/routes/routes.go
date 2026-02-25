package routes

import (
	"net/http"

	"hotel-booking/internal/handlers"
	"hotel-booking/internal/middleware"
)

func Register(mux *http.ServeMux, app handlers.AppContext) {
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/hotel", app.HotelDetail)
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			app.LoginPost(w, r)
			return
		}
		app.LoginPage(w, r)
	})
	mux.HandleFunc("/booking", middleware.Auth(app.BookingPage))
	mux.HandleFunc("/history", middleware.Auth(app.HistoryPage))
	mux.HandleFunc("/payment", middleware.Auth(app.PaymentPage))
	mux.HandleFunc("/admin", middleware.Auth(middleware.Admin(app.AdminDashboard)))
}
