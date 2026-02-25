package database

import "hotel-booking/internal/models"

type DB struct {
	Users  []models.User
	Hotels []models.Hotel
	Rooms  []models.Room
}

func New() *DB {
	return &DB{}
}
