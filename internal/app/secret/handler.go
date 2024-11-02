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
	// Check if user is authorized
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	// Read the JSON payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Validate and decode payload
	var passwordReq model.GeneratePasswordRequest
	if err := crypto.ValidatePayload(data, &passwordReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Generate the password
	passwordGenerated, err := h.service.GeneratePassword(r.Context(), passwordReq.Length, passwordReq.IncludeSpecialSymbol)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Respond with the generated password
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"password": passwordGenerated,
	})
}

// Create handles the creation of a new secret
func (h *SecretHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Check if the user is authorized
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	// Read the JSON payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Decode payload into the Secret struct
	var secret model.Secret
	if err := json.Unmarshal(data, &secret); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Assign the user ID from the authenticated user
	secret.UserID = user.ID

	// Check if essential fields are missing
	if secret.Title == "" || secret.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Encrypt the password before storing
	secret.Password, err = crypto.Encrypt(secret.Password, []byte(h.config.EncryptionKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Set timestamps
	secret.CreatedAt = time.Now()
	secret.UpdatedAt = &secret.CreatedAt

	// Call the service to create the secret
	if err := h.service.Create(r.Context(), secret); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "secret created successfully",
	})
}

func (h *SecretHandler) List(w http.ResponseWriter, r *http.Request) {
	// Check if user is authorized
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	// Fetch the list of secrets for the user
	secrets, err := h.service.List(r.Context(), user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Decrypt each secret's password
	for i := range secrets {
		decryptedPassword, err := crypto.Decrypt(secrets[i].Password, []byte(h.config.EncryptionKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		secrets[i].Password = decryptedPassword
	}

	// Respond with the list of decrypted secrets
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secrets)
}
