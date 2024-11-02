package secret

import (
	"encoding/json"
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

func (h *SecretHandler) GeneratePassword(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	// Read and validate payload
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
	if err := json.Unmarshal(data, &passwordReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	fmt.Printf("Parsed Length: %d, IncludeSpecialSymbol: %v\n", passwordReq.Length, passwordReq.IncludeSpecialSymbol)

	// Generate password
	passwordGenerated, err := h.service.GeneratePassword(r.Context(), passwordReq.Length, passwordReq.IncludeSpecialSymbol)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
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
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	// Read and decode payload
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var secret model.Secret
	if err := json.Unmarshal(data, &secret); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Set user ID and check fields
	secret.UserID = user.ID
	if secret.Title == "" || secret.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Title and password cannot be empty",
		})
		return
	}

	// Encrypt password and create the secret
	secret.Password, err = crypto.Encrypt(secret.Password, []byte(h.config.EncryptionKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	secret.CreatedAt = time.Now()
	secret.UpdatedAt = &secret.CreatedAt
	if err := h.service.Create(r.Context(), secret); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
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
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	// Fetch secrets and decrypt passwords
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

	// Respond with SuccessResponse
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Secrets retrieved successfully",
		Data:    secrets,
	})
}
