package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct{ base string }

func NewRenderer(base string) *Renderer { return &Renderer{base: base} }

func (r *Renderer) Render(w http.ResponseWriter, page string, data any) {
	layout := filepath.Join(r.base, "layouts", "base.html")
	tpl, err := template.ParseFiles(layout, filepath.Join(r.base, page))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_ = tpl.ExecuteTemplate(w, "base", data)
}
