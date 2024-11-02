package secret

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/auth"
	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/internal/utils/database"
	"github.com/mahinops/secretcli-web/model"
)

type SecretHandler struct {
	service  *SecretService
	config   *database.Config
	renderer *tmplrndr.Renderer
}

// NewSecretHandler creates a new instance of SecretHandler// NewSecretHandler creates a new instance of SecretHandler
func NewSecretHandler(service *SecretService, config *database.Config, renderer *tmplrndr.Renderer) *SecretHandler {
	return &SecretHandler{service: service, config: config, renderer: renderer} // Update this line
}

func (h *SecretHandler) SecretCreateForm(w http.ResponseWriter, r *http.Request) {
	if h.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	h.renderer.Render(w, "secrets.create.form", nil)
}

func (h *SecretHandler) SecretListTemplate(w http.ResponseWriter, r *http.Request) {
	if h.renderer == nil {
		http.Error(w, "Renderer is not initialized", http.StatusInternalServerError)
		return
	}
	h.renderer.Render(w, "secrets.table", nil)
}

func (h *SecretHandler) GeneratePassword(w http.ResponseWriter, r *http.Request) {
	// Check If User is Authorized
	user, err := auth.ValidateToken(r) // This function should return the user if authenticated
	if err != nil || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Read the JSON payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	// Validate the payload against GeneratePasswordRequest struct
	var passwordReq model.GeneratePasswordRequest
	if err := crypto.ValidatePayload(data, &passwordReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Decode JSON into the struct after validation
	if err := json.Unmarshal(data, &passwordReq); err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	// Generate the password using the crypto package
	passwordGenerated, err := h.service.GeneratePassword(r.Context(), passwordReq.Length, passwordReq.IncludeSpecialSymbol)
	if err != nil {
		http.Error(w, "Error generating password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return Generated Password
	w.Write([]byte(passwordGenerated))
}

// Create handles the creation of a new secret
func (h *SecretHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Check If User is Authorized
	user, err := auth.ValidateToken(r) // This function should return the user if authenticated
	if err != nil || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Read the JSON payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	// Validate the payload against GeneratePasswordRequest struct
	var secret model.Secret
	if err := crypto.ValidatePayload(data, &secret); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Decode JSON into the struct after validation
	if err := json.Unmarshal(data, &secret); err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
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

	// Decrypt the passwords in the fetched secrets
	for i := range secrets {
		decryptedPassword, err := crypto.Decrypt(secrets[i].Password, []byte(h.config.EncryptionKey)) // Decrypt the password
		if err != nil {
			http.Error(w, "Error decrypting password: "+err.Error(), http.StatusInternalServerError)
			return
		}
		secrets[i].Password = decryptedPassword // Replace the encrypted password with the decrypted one
	}

	// Encode the response as JSON
	w.WriteHeader(http.StatusOK) // Set the status to 200 OK
	json.NewEncoder(w).Encode(secrets)
}
