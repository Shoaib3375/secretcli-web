package secret

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mahinops/secretcli-web/model"
	"github.com/mahinops/secretcli-web/utils/auth"
)

type SecretHandler struct {
	service *SecretService
}

// NewSecretHandler creates a new instance of SecretHandler
func NewSecretHandler(service *SecretService) *SecretHandler {
	return &SecretHandler{service: service}
}

// Create handles the creation of a new secret
func (h *SecretHandler) Create(w http.ResponseWriter, r *http.Request) {

	user, err := auth.ValidateToken(r) // This function should return the user if authenticated
	if err != nil || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var secret model.Secret
	if err := json.NewDecoder(r.Body).Decode(&secret); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Set timestamps
	secret.CreatedAt = time.Now()
	secret.UpdatedAt = &secret.CreatedAt

	// Call the service to create the secret
	if err := h.service.Create(r.Context(), secret); err != nil {
		http.Error(w, "Error creating secret: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // Set the status to 201 Created
	response := map[string]string{
		"message": "Secret created successfully",
	}
	json.NewEncoder(w).Encode(response) // Encode the response as JSON
}
