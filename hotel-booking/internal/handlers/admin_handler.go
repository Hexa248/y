package handlers

import "net/http"

type AdminHandler struct{ renderer *Renderer }

func NewAdminHandler(r *Renderer) *AdminHandler { return &AdminHandler{renderer: r} }

func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "admin/dashboard.html", map[string]any{"Title": "Admin Dashboard"})
}
