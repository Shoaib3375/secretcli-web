// templates/renderer.go
package tmplrndr

import (
	"html/template"
	"log"
	"net/http"
)

type Renderer struct {
	templates *template.Template
}

// NewRenderer initializes the template renderer by parsing all templates in the folder.
func NewRenderer(pattern string) *Renderer {
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}
	return &Renderer{templates: tmpl}
}

// Render renders a specific template by name
func (r *Renderer) Render(w http.ResponseWriter, name string, data interface{}) {
	err := r.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
