package models

type City struct {
	ID        int
	Name      string
	ImageURL  string
	Latitude  float64
	Longitude float64
}

type Hotel struct {
	ID          int
	CityID      int
	Name        string
	Address     string
	Description string
	ImageURL    string
	Rating      float64
	Latitude    float64
	Longitude   float64
	Policies    []string
	Facilities  []string
	Reviews     []Review
}

type Review struct {
	Author  string
	Rating  float64
	Comment string
}
