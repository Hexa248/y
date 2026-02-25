package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"hotel-booking/internal/models"
)

type cityShowcase struct {
	Name     string
	ImageURL string
	Count    int
	Tagline  string
}

type cityPageData struct {
	Name       string
	ImageURL   string
	HotelCount int
	Hotels     []models.Hotel
}

func (a AppContext) Home(w http.ResponseWriter, r *http.Request) {
	hotels := a.Hotels.HotelsWithRooms()

	cityMap := map[string]cityShowcase{}
	for _, h := range hotels {
		if existing, ok := cityMap[h.City]; ok {
			existing.Count++
			cityMap[h.City] = existing
		} else {
			cityMap[h.City] = cityShowcase{
				Name:     h.City,
				ImageURL: h.ImageURL,
				Count:    1,
				Tagline:  cityTagline(h.City),
			}
		}
	}

	showcases := make([]cityShowcase, 0, len(cityMap))
	for _, city := range []string{"Jakarta", "Bandung", "Yogyakarta", "Surabaya", "Bali", "Medan", "Makassar", "Semarang", "Lombok", "Malang", "Bogor", "Batam", "Padang", "Palembang", "Manado", "Pontianak", "Banjarmasin", "Balikpapan", "Pekanbaru", "Solo"} {
		if c, ok := cityMap[city]; ok {
			showcases = append(showcases, c)
		}
	}

	_ = a.Templates.ExecuteTemplate(w, "index.html", map[string]any{
		"CityShowcases": showcases,
		"TotalHotels":   len(hotels),
	})
}

func (a AppContext) CityHotels(w http.ResponseWriter, r *http.Request) {
	cityName := strings.TrimSpace(r.URL.Query().Get("name"))
	if cityName == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	hotels := a.Hotels.HotelsWithRooms()
	cityHotels := make([]models.Hotel, 0, 10)
	cityImage := ""
	for _, h := range hotels {
		if strings.EqualFold(h.City, cityName) {
			cityHotels = append(cityHotels, h)
			if cityImage == "" {
				cityImage = h.ImageURL
			}
		}
	}

	if len(cityHotels) == 0 {
		http.NotFound(w, r)
		return
	}

	sort.Slice(cityHotels, func(i, j int) bool {
		return cityHotels[i].ID < cityHotels[j].ID
	})

	_ = a.Templates.ExecuteTemplate(w, "city.html", cityPageData{
		Name:       cityHotels[0].City,
		ImageURL:   cityImage,
		HotelCount: len(cityHotels),
		Hotels:     cityHotels,
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

func cityTagline(city string) string {
	taglines := map[string]string{
		"Jakarta":     "Skyline metropolitan & hiburan premium",
		"Bandung":     "Udara sejuk, kuliner, dan kreativitas",
		"Yogyakarta":  "Warisan budaya, seni, dan keramahan",
		"Surabaya":    "Kota bisnis modern dengan sejarah kuat",
		"Bali":        "Pulau resort tropis kelas dunia",
		"Medan":       "Gerbang Sumatra dengan wisata kuliner",
		"Makassar":    "City escape bahari di timur Indonesia",
		"Semarang":    "Kota lama cantik dan wisata keluarga",
		"Lombok":      "Pantai eksotis dan petualangan alam",
		"Malang":      "Destinasi pegunungan yang adem",
		"Bogor":       "Staycation hijau dekat Jabodetabek",
		"Batam":       "Akses cepat untuk perjalanan bisnis",
		"Padang":      "Kuliner otentik dan panorama pantai",
		"Palembang":   "Kota sungai dengan sejarah Sriwijaya",
		"Manado":      "Terumbu karang dan wisata laut terbaik",
		"Pontianak":   "Kota khatulistiwa dengan pesona unik",
		"Banjarmasin": "Jelajah pasar terapung legendaris",
		"Balikpapan":  "Kota modern yang rapi dan nyaman",
		"Pekanbaru":   "Transit strategis untuk Sumatra tengah",
		"Solo":        "Nuansa budaya Jawa yang elegan",
	}
	if tagline, ok := taglines[city]; ok {
		return tagline
	}
	return "Temukan pengalaman menginap terbaik"
}
