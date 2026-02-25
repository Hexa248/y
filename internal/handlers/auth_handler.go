package handlers

import "net/http"

func (a AppContext) LoginPage(w http.ResponseWriter, r *http.Request) {
	_ = a.Templates.ExecuteTemplate(w, "login.html", nil)
}

func (a AppContext) LoginPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	u, token, ok := a.Auth.Login(r.FormValue("email"), r.FormValue("password"))
	if !ok {
		_ = a.Templates.ExecuteTemplate(w, "login.html", map[string]any{"Error": "Email / password salah"})
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "token", Value: token, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "role", Value: string(u.Role), Path: "/"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
