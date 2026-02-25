package handlers

import "net/http"

func (a AppContext) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	_ = a.Templates.ExecuteTemplate(w, "dashboard.html", nil)
}
