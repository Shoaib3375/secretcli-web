package auth

import (
	"net/http"

	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
)

type AuthWebHandler struct {
	renderer *tmplrndr.Renderer
}

func NewAuthWebHandler(renderer *tmplrndr.Renderer) *AuthWebHandler {
	return &AuthWebHandler{renderer: renderer}
}

func (web *AuthWebHandler) LoginUserForm(w http.ResponseWriter, r *http.Request) {
	if web.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	web.renderer.Render(w, "auth.login.form", nil)
}

func (web *AuthWebHandler) RegisterUserForm(w http.ResponseWriter, r *http.Request) {
	if web.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	web.renderer.Render(w, "auth.registration.form", nil)
}
