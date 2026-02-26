package repositories

import (
	"fmt"

	"hotel-booking/internal/models"
)

type HotelRepository struct {
	cities []models.City
	hotels []models.Hotel
	rooms  []models.Room
}

func NewHotelRepository() *HotelRepository {
	r := &HotelRepository{}
	r.seed()
	return r
}

func (r *HotelRepository) Cities() []models.City  { return r.cities }
func (r *HotelRepository) Hotels() []models.Hotel { return r.hotels }

func (r *HotelRepository) HotelsByCity(cityID int) []models.Hotel {
	out := []models.Hotel{}
	for _, h := range r.hotels {
		if h.CityID == cityID {
			out = append(out, h)
		}
	}
	return out
}

func (r *HotelRepository) HotelByID(id int) (models.Hotel, bool) {
	for _, h := range r.hotels {
		if h.ID == id {
			return h, true
		}
	}
	return models.Hotel{}, false
}

func (r *HotelRepository) RoomsByHotel(hotelID int) []models.Room {
	out := []models.Room{}
	for _, rm := range r.rooms {
		if rm.HotelID == hotelID {
			out = append(out, rm)
		}
	}
	return out
}

func (r *HotelRepository) RoomByID(id int) (models.Room, bool) {
	for _, rm := range r.rooms {
		if rm.ID == id {
			return rm, true
		}
	}
	return models.Room{}, false
}

func (r *HotelRepository) seed() {
	cities := []string{"Jakarta", "Bandung", "Surabaya", "Yogyakarta", "Bali", "Lombok", "Malang", "Semarang", "Medan", "Palembang", "Makassar", "Balikpapan", "Samarinda", "Manado", "Batam", "Pekanbaru", "Padang", "Banjarmasin", "Jayapura", "Labuan Bajo"}
	hotelID := 1
	roomID := 1
	for i, c := range cities {
		cityID := i + 1
		r.cities = append(r.cities, models.City{ID: cityID, Name: c, ImageURL: fmt.Sprintf("https://source.unsplash.com/featured/800x500/?%s,indonesia,landmark", c), Latitude: -6.2 + float64(i)/10, Longitude: 106.8 + float64(i)/10})
		for h := 1; h <= 10; h++ {
			h := models.Hotel{ID: hotelID, CityID: cityID, Name: fmt.Sprintf("%s Signature Hotel %d", c, h), Address: fmt.Sprintf("Jl. Ikonik %d, %s", h, c), Description: "Hotel modern dengan sentuhan desain travel marketplace premium.", ImageURL: fmt.Sprintf("https://source.unsplash.com/featured/900x600/?hotel,%s,indonesia", c), Rating: 4.0 + float64(h%10)/10, Latitude: -6.2 + float64(h)/100, Longitude: 106.8 + float64(h)/100, Policies: []string{"Check-in 14:00", "Check-out 12:00", "Non-smoking area"}, Facilities: []string{"WiFi", "Kolam renang", "Gym", "Sarapan"}, Reviews: []models.Review{{Author: "Rina", Rating: 4.8, Comment: "Lokasi strategis dan kamar bersih."}, {Author: "Andi", Rating: 4.6, Comment: "Pelayanan ramah, cocok keluarga."}}}
			r.hotels = append(r.hotels, h)
			for rm := 1; rm <= 10; rm++ {
				rt := models.RoomStandard
				fac := []string{"AC", "WiFi", "TV"}
				price := 450000 + rm*50000
				if rm > 3 && rm <= 7 {
					rt = models.RoomDeluxe
					fac = append(fac, "Bathtub", "City View")
					price += 300000
				}
				if rm > 7 {
					rt = models.RoomVIP
					fac = append(fac, "Private Lounge", "Butler", "Premium Amenities")
					price += 700000
				}
				r.rooms = append(r.rooms, models.Room{ID: roomID, HotelID: hotelID, Type: rt, Name: fmt.Sprintf("%s Room %d", rt, rm), Price: price, Beds: 1 + rm%2, Capacity: 2 + rm%3, SizeM2: 24 + rm*2, Facilities: fac, FloorPlan: fmt.Sprintf("Denah kamar tipe %s #%d", rt, rm), Images: []string{fmt.Sprintf("https://source.unsplash.com/featured/640x420/?hotel-room,%s", rt)}, MapEmbedURL: "https://maps.google.com/maps?q=-6.2,106.8&z=14&output=embed"})
				roomID++
			}
			hotelID++
		}
	}
}
