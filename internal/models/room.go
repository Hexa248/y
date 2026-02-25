package models

type Room struct {
	ID          int
	HotelID     int
	Name        string
	Price       int
	Beds        int
	Capacity    int
	Facilities  []string
	FloorPlan   string
	Description string
}
