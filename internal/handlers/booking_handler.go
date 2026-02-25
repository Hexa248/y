package handlers

import "net/http"

func (a AppContext) BookingPage(w http.ResponseWriter, r *http.Request) {
	_ = a.Templates.ExecuteTemplate(w, "booking.html", nil)
}

func (a AppContext) HistoryPage(w http.ResponseWriter, r *http.Request) {
	_ = a.Templates.ExecuteTemplate(w, "history.html", nil)
}
