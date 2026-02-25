package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"hotel-booking/config"
	"hotel-booking/database"
	"hotel-booking/database/seed"
	"hotel-booking/internal/handlers"
	"hotel-booking/internal/repositories"
	"hotel-booking/internal/routes"
	"hotel-booking/internal/services"
)

func main() {
	cfg := config.Load()
	db := database.New()
	seed.Load(db)
	t := template.Must(template.ParseGlob("views/**/*.html"))

	app := handlers.AppContext{
		Templates: t,
		Hotels:    services.HotelService{Hotels: repositories.HotelRepository{DB: db}, Rooms: repositories.RoomRepository{DB: db}},
		Auth:      services.AuthService{Users: repositories.UserRepository{DB: db}},
	}

	mux := http.NewServeMux()
	routes.Register(mux, app)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("%s running at http://localhost:%s", cfg.AppName, cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux))
}
