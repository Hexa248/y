package handlers

import "net/http"

func (a AppContext) PaymentPage(w http.ResponseWriter, r *http.Request) {
	_ = a.Templates.ExecuteTemplate(w, "payment.html", nil)
}
