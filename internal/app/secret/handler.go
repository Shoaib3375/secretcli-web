package secret

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mahinops/secretcli-web/internal/utils/auth"
	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/internal/utils/database"
	"github.com/mahinops/secretcli-web/model"
)

type SecretHandler struct {
	service *SecretService
	config  *database.Config
}

// NewSecretHandler creates a new instance of SecretHandler// NewSecretHandler creates a new instance of SecretHandler
func NewSecretHandler(service *SecretService, config *database.Config) *SecretHandler {
	return &SecretHandler{service: service, config: config} // Update this line
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

	// Set UserID from the authenticated user
	secret.UserID = user.ID

	// Encrypt the password before storing it
	secret.Password, err = crypto.Encrypt(secret.Password, []byte(h.config.EncryptionKey)) // Use the key from config
	if err != nil {
		http.Error(w, "Error encrypting password: "+err.Error(), http.StatusInternalServerError)
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

func (h *SecretHandler) List(w http.ResponseWriter, r *http.Request) {

	user, err := auth.ValidateToken(r) // This function should return the user if authenticated
	if err != nil || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	secrets, err := h.service.List(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Error fetching secrets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response as JSON
	w.WriteHeader(http.StatusOK) // Set the status to 200 OK
	json.NewEncoder(w).Encode(secrets)
}
