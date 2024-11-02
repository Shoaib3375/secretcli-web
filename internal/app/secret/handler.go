package secret

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (h *SecretHandler) handleError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Code:    code,
		Message: err.Error(),
	})
}

func (h *SecretHandler) GeneratePassword(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	// Read and validate payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Validate and decode payload
	var passwordReq model.GeneratePasswordRequest
	if err := crypto.ValidatePayload(data, &passwordReq); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Decode JSON into the struct after validation
	if err := json.Unmarshal(data, &passwordReq); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Generate password
	passwordGenerated, err := h.service.GeneratePassword(r.Context(), passwordReq.Length, passwordReq.IncludeSpecialSymbol)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with SuccessResponse
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Password generated successfully",
		Data:    map[string]string{"password": passwordGenerated},
	})
}

// Create handles the creation of a new secret
func (h *SecretHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	// Read and decode payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	var secret model.Secret
	if err := crypto.ValidatePayload(data, &secret); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	if err := json.Unmarshal(data, &secret); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Set user ID and check fields
	secret.UserID = user.ID
	if secret.Title == "" || secret.Password == "" {
		h.handleError(w, http.StatusBadRequest, errors.New("title and password cannot be empty"))
		return
	}

	// Encrypt password and create the secret
	secret.Password, err = crypto.Encrypt(secret.Password, []byte(h.config.EncryptionKey))
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	secret.CreatedAt = time.Now()
	secret.UpdatedAt = &secret.CreatedAt
	if err := h.service.Create(r.Context(), secret); err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with SuccessResponse
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "Secret created successfully",
	})
}

func (h *SecretHandler) List(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}
	fmt.Println("user", user.ID, user.Name)

	// Fetch secrets and decrypt passwords
	secrets, err := h.service.List(r.Context(), user.ID)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Decrypt each secret's password
	for i := range secrets {
		// Validate Base64 string
		if !crypto.IsValidBase64(secrets[i].Password) {
			h.handleError(w, http.StatusBadRequest, fmt.Errorf("invalid Base64 string"))
			return
		}

		decryptedPassword, err := crypto.Decrypt(secrets[i].Password, []byte(h.config.EncryptionKey))
		if err != nil {
			h.handleError(w, http.StatusInternalServerError, err)
			return
		}
		secrets[i].Password = decryptedPassword
	}

	// Respond with SuccessResponse
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Secrets retrieved successfully",
		Data:    secrets,
	})
}
