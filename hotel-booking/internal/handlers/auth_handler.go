package handlers

import (
	"net/http"

	"hotel-booking/internal/services"
	"hotel-booking/pkg/security"
)

type AuthHandler struct {
	renderer *Renderer
	auth     *services.AuthService
}

func NewAuthHandler(r *Renderer, a *services.AuthService) *AuthHandler {
	return &AuthHandler{renderer: r, auth: a}
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "auth/login.html", map[string]any{"Title": "Login"})
}
func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "auth/register.html", map[string]any{"Title": "Register"})
}

func (h *AuthHandler) DoLogin(w http.ResponseWriter, r *http.Request) {
	email, password := r.FormValue("email"), r.FormValue("password")
	u, err := h.auth.Login(email, password)
	if err != nil {
		http.Redirect(w, r, "/login?err=1", http.StatusSeeOther)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "session", Value: security.GenerateToken(u.ID, string(u.Role)), Path: "/"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AuthHandler) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login?msg=google_oauth_demo_only", http.StatusSeeOther)
}
