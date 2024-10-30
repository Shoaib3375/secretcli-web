package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// LoginUserSubmit receives form data and forwards it to LoginUser
func (h *AuthHandler) LoginUserSubmit(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create a user object from the form data
	user := model.Auth{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Convert user data to JSON
	userData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user data", http.StatusInternalServerError)
		return
	}
	resp, err := http.Post("http://localhost:8080/auth/api/login", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Read the response body
		errorMessage := fmt.Sprintf(string(body))
		data := struct {
			Error string
		}{
			Error: errorMessage,
		}
		h.renderer.Render(w, "auth.login.form", data) // Render form with error
		return
	}

	// Decode the response body for the success case
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	http.Redirect(w, r, "/auth/web/register", http.StatusSeeOther)
}

// RegisterUserSubmit receives form data and forwards it to RegisterUser
func (h *AuthHandler) RegisterUserSubmit(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create a user object from the form data
	user := model.Auth{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Convert user data to JSON
	userData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user data", http.StatusInternalServerError)
		return
	}
	resp, err := http.Post("http://localhost:8080/auth/api/register", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body) // Read the response body
		errorMessage := fmt.Sprintf(string(body))
		data := struct {
			Error string
		}{
			Error: errorMessage,
		}
		h.renderer.Render(w, "auth.registration.form", data)
		return
	}
	http.Redirect(w, r, "/auth/web/login", http.StatusSeeOther)
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
