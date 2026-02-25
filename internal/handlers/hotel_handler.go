package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"hotel-booking/internal/models"
)

type cityShowcase struct {
	Name     string
	ImageURL string
	Count    int
}

func (a AppContext) Home(w http.ResponseWriter, r *http.Request) {
	selectedCity := strings.TrimSpace(r.URL.Query().Get("city"))
	hotels := a.Hotels.HotelsWithRooms()

	cityMap := map[string]cityShowcase{}
	filtered := make([]models.Hotel, 0)
	for _, h := range hotels {
		if existing, ok := cityMap[h.City]; ok {
			existing.Count++
			cityMap[h.City] = existing
		} else {
			cityMap[h.City] = cityShowcase{Name: h.City, ImageURL: h.ImageURL, Count: 1}
		}
		if selectedCity == "" || strings.EqualFold(selectedCity, h.City) {
			filtered = append(filtered, h)
		}
	}

	showcases := make([]cityShowcase, 0, len(cityMap))
	for _, city := range []string{"Jakarta", "Bandung", "Yogyakarta", "Surabaya", "Bali", "Medan", "Makassar", "Semarang", "Lombok", "Malang", "Bogor", "Batam", "Padang", "Palembang", "Manado", "Pontianak", "Banjarmasin", "Balikpapan", "Pekanbaru", "Solo"} {
		if c, ok := cityMap[city]; ok {
			showcases = append(showcases, c)
		}
	}

	_ = a.Templates.ExecuteTemplate(w, "index.html", map[string]any{
		"Hotels":        filtered,
		"CityShowcases": showcases,
		"SelectedCity":  selectedCity,
	})
}

func (a AppContext) HotelDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	hotel, rooms, ok := a.Hotels.HotelDetail(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	_ = a.Templates.ExecuteTemplate(w, "detail.html", map[string]any{"Hotel": hotel, "Rooms": rooms})
}
