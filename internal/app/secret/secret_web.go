package secret

import (
	"net/http"

	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
)

type SecretWebHandler struct {
	renderer *tmplrndr.Renderer
}

func NewSecretWebHandler(renderer *tmplrndr.Renderer) *SecretWebHandler {
	return &SecretWebHandler{renderer: renderer}
}

func (web *SecretWebHandler) SecretCreateForm(w http.ResponseWriter, r *http.Request) {
	if web.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	web.renderer.Render(w, "secrets.create.form", nil)
}

func (web *SecretWebHandler) SecretListTemplate(w http.ResponseWriter, r *http.Request) {
	if web.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	web.renderer.Render(w, "secrets.table", nil)
}
