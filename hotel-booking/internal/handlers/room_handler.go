package handlers

import (
	"net/http"

	"hotel-booking/internal/services"
	"hotel-booking/pkg/utils"
)

type RoomHandler struct {
	renderer *Renderer
	hotel    *services.HotelService
}

func NewRoomHandler(r *Renderer, h *services.HotelService) *RoomHandler {
	return &RoomHandler{renderer: r, hotel: h}
}

func (h *RoomHandler) Detail(w http.ResponseWriter, r *http.Request) {
	id := utils.MustAtoi(r.URL.Query().Get("id"))
	room, ok := h.hotel.RoomByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	h.renderer.Render(w, "hotels/room.html", map[string]any{"Title": room.Name, "Room": room})
}
