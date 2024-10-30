package auth

import (
	"encoding/json"
	"net/http"

	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/auth"
	"github.com/mahinops/secretcli-web/model"
)

type AuthHandler struct {
	usecase  model.AuthUsecase
	renderer *tmplrndr.Renderer
}

func NewAuthHandler(usecase model.AuthUsecase, renderer *tmplrndr.Renderer) *AuthHandler {
	return &AuthHandler{usecase: usecase, renderer: renderer}
}

// RegisterUserForm renders the registration form
func (h *AuthHandler) LoginUserForm(w http.ResponseWriter, r *http.Request) {
	if h.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	h.renderer.Render(w, "auth.login.form", nil)
}

// RegisterUserForm renders the registration form
func (h *AuthHandler) RegisterUserForm(w http.ResponseWriter, r *http.Request) {
	if h.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	h.renderer.Render(w, "auth.registration.form", nil)
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.Auth
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	name, err := h.usecase.Create(r.Context(), user)
	if err != nil {
		if err.Error() == "email already exists" {
			http.Error(w, err.Error(), http.StatusConflict) // 409 Conflict for existing email
			return
		}
		http.Error(w, "Error registering user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message and the user's name
	w.WriteHeader(http.StatusCreated) // Set the status to 201 Created
	response := map[string]string{
		"message": "User " + name + " created successfully",
	}
	json.NewEncoder(w).Encode(response) // Encode the response as JSON
}

// UserLogin is the structure for the login request body
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest UserLogin
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Use the login service to authenticate the user
	user, err := h.usecase.Login(r.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate a JWT token for the user
	token, err := auth.GenerateToken(user) // Reference the generateToken from utils/auth/jwt.go
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Respond with the generated token and expiry
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":  token,
		"expiry": user.Expiry, // Include the expiry time in the response
	})
}
