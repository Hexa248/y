package handlers

import (
	"hotel-booking/internal/services"
	"html/template"
)

type AppContext struct {
	Templates *template.Template
	Hotels    services.HotelService
	Auth      services.AuthService
}
