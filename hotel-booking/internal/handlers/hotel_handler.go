package handlers

import (
	"net/http"

	"hotel-booking/internal/services"
	"hotel-booking/pkg/utils"
)

type HotelHandler struct {
	renderer *Renderer
	hotel    *services.HotelService
}

func NewHotelHandler(r *Renderer, h *services.HotelService) *HotelHandler {
	return &HotelHandler{renderer: r, hotel: h}
}

func (h *HotelHandler) Home(w http.ResponseWriter, r *http.Request) {
	cities := h.hotel.Cities()
	cityNames := map[int]string{}
	for _, city := range cities {
		cityNames[city.ID] = city.Name
	}
	h.renderer.Render(w, "hotels/index.html", map[string]any{
		"Title":     "NusaStay",
		"Cities":    cities,
		"Hotels":    h.hotel.Hotels(),
		"CityNames": cityNames,
	})
}

func (h *HotelHandler) CityHotels(w http.ResponseWriter, r *http.Request) {
	cityID := utils.MustAtoi(r.URL.Query().Get("city"))
	h.renderer.Render(w, "hotels/city.html", map[string]any{"Title": "Daftar Hotel", "Hotels": h.hotel.HotelsByCity(cityID), "CityID": cityID})
}

func (h *HotelHandler) Detail(w http.ResponseWriter, r *http.Request) {
	id := utils.MustAtoi(r.URL.Query().Get("id"))
	hotel, ok := h.hotel.HotelByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	h.renderer.Render(w, "hotels/detail.html", map[string]any{"Title": hotel.Name, "Hotel": hotel, "Rooms": h.hotel.RoomsByHotel(id)})
}
