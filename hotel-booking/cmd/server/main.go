package main

import (
	"log"
	"net/http"

	"hotel-booking/config"
	"hotel-booking/internal/handlers"
	"hotel-booking/internal/middleware"
	"hotel-booking/internal/repositories"
	"hotel-booking/internal/routes"
	"hotel-booking/internal/services"
)

func main() {
	_ = config.LoadEnv(".env")
	cfg := config.Load()

	render := handlers.NewRenderer("views")
	hotelRepo := repositories.NewHotelRepository()
	userRepo := repositories.NewUserRepository()
	bookingRepo := repositories.NewBookingRepository()
	paymentRepo := repositories.NewPaymentRepository()

	authSvc := services.NewAuthService(userRepo)
	hotelSvc := services.NewHotelService(hotelRepo)
	bookingSvc := services.NewBookingService(bookingRepo)
	paymentSvc := services.NewPaymentService(paymentRepo)

	authH := handlers.NewAuthHandler(render, authSvc)
	hotelH := handlers.NewHotelHandler(render, hotelSvc)
	roomH := handlers.NewRoomHandler(render, hotelSvc)
	bookingH := handlers.NewBookingHandler(render, bookingSvc)
	paymentH := handlers.NewPaymentHandler(render, paymentSvc)
	adminH := handlers.NewAdminHandler(render)

	mux := http.NewServeMux()
	routes.Register(mux, authH, hotelH, roomH, bookingH, paymentH, adminH)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("%s running at :%s", cfg.AppName, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, middleware.Logging(mux)); err != nil {
		log.Fatal(err)
	}
}
