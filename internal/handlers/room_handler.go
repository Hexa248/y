package handlers

import "net/http"

func (a AppContext) RoomPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
