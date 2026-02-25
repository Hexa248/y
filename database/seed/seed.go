package seed

import (
	"fmt"
	"hotel-booking/database"
	"hotel-booking/internal/models"
)

func Load(db *database.DB) {
	cities := []string{"Jakarta", "Bandung", "Yogyakarta", "Surabaya", "Bali", "Medan", "Makassar", "Semarang", "Lombok", "Malang", "Bogor", "Batam", "Padang", "Palembang", "Manado", "Pontianak", "Banjarmasin", "Balikpapan", "Pekanbaru", "Solo"}
	baseLat := -6.2
	baseLng := 106.8

	db.Users = []models.User{
		{ID: 1, Name: "Admin", Email: "admin@nusstay.com", Password: "admin123", Role: models.RoleAdmin},
		{ID: 2, Name: "Guest", Email: "guest@nusstay.com", Password: "guest123", Role: models.RoleUser},
		{ID: 3, Name: "Staff", Email: "staff@nusstay.com", Password: "staff123", Role: models.RoleHotelStaff},
	}

	hID, rID := 1, 1
	for cIdx, city := range cities {
		for i := 1; i <= 10; i++ {
			lat := baseLat + float64(cIdx)*0.35 + float64(i)*0.01
			lng := baseLng + float64(cIdx)*0.4 + float64(i)*0.01
			hotel := models.Hotel{
				ID:          hID,
				City:        city,
				Name:        fmt.Sprintf("%s Grand Stay %d", city, i),
				Description: "Design modern dengan perpaduan gaya Traveloka, Booking, tiket.com, dan Trip.com.",
				ImageURL:    fmt.Sprintf("https://source.unsplash.com/featured/800x500/?%s,landmark", city),
				MapURL:      fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&z=14&output=embed", lat, lng),
				Address:     fmt.Sprintf("Jl. Ikonik %d, %s", i, city),
				Facilities:  []string{"WiFi 500Mbps", "Kolam renang", "Gym", "Sarapan", "Airport shuttle"},
				Comments:    []models.Comment{{User: "Ayu", Message: "Hotel bersih dan strategis", Rating: 5}, {User: "Bima", Message: "Pelayanan ramah", Rating: 4}},
			}
			db.Hotels = append(db.Hotels, hotel)
			for j := 1; j <= 10; j++ {
				db.Rooms = append(db.Rooms, models.Room{
					ID:          rID,
					HotelID:     hID,
					Name:        fmt.Sprintf("Tipe Kamar %d", j),
					Price:       350000 + (j * 120000),
					Beds:        1 + j%3,
					Capacity:    2 + j%4,
					Facilities:  []string{"AC", "Smart TV", "Bathtub", "Coffee maker", "Balkon"},
					FloorPlan:   fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&z=19&output=embed", lat+0.002*float64(j), lng+0.001*float64(j)),
					Description: "Kamar stylish dengan pencahayaan hangat dan interior modern.",
				})
				rID++
			}
			hID++
		}
	}
}
