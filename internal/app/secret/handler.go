package secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"

	"github.com/mahinops/secretcli-web/internal/utils/auth"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/model"
)

type SecretHandler struct {
	service      *SecretService
	commonConfig *common.CommonConfig
}

// NewSecretHandler creates a new instance of SecretHandler// NewSecretHandler creates a new instance of SecretHandler
func NewSecretHandler(service *SecretService, commonConfig *common.CommonConfig) *SecretHandler {
	return &SecretHandler{service: service, commonConfig: commonConfig} // Update this line
}

func (h *SecretHandler) handleError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Code:    code,
		Message: err.Error(),
	})
}

// GeneratePassword handles the generation of a password
//
//	@Summary		Generate a secure password
//	@Description	Generates a password based on provided parameters.
//	@Tags			secrets
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string									true	"Bearer <token>"
//	@Param			body			body		model.SwaggerGeneratePasswordRequest	true	"Password generation parameters"
//	@Success		200				{object}	model.SuccessResponse
//	@Failure		400				{object}	model.ErrorResponse
//	@Failure		401				{object}	model.ErrorResponse
//	@Failure		500				{object}	model.ErrorResponse
//	@Router			/secret/api/generatepassword [post]
func (h *SecretHandler) GeneratePassword(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r, h.commonConfig.JWTSecretKey)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	var passwordReq model.GeneratePasswordRequest
	if err := common.ParseAndValidatePayload(r, &passwordReq); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	passwordGenerated, err := h.service.GeneratePassword(r.Context(), passwordReq.Length, passwordReq.IncludeSpecialSymbol)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	common.RespondWithSuccess(w, http.StatusOK, "Password generated successfully", map[string]interface{}{
		"password": passwordGenerated,
	})
}

// Create handles the creation of a new secret.
//
//	@Summary		Create a new secret
//	@Description	This endpoint allows users to create a new secret with an encrypted password.
//	@Tags			secrets
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token for authentication"
//	@Param			secret			body		model.SwaggerSecretRequest	true	"Secret payload containing title, username, password, note, email, website, and user ID"
//	@Success		201				{object}	model.SuccessResponse
//	@Failure		400				{object}	model.ErrorResponse
//	@Failure		401				{object}	model.ErrorResponse
//	@Failure		500				{object}	model.ErrorResponse
//	@Router			/secret/api/create [post]
func (h *SecretHandler) Create(w http.ResponseWriter, r *http.Request) {
	user, err := auth.ValidateToken(r, h.commonConfig.JWTSecretKey)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	var secret model.Secret
	if err := common.ParseAndValidatePayload(r, &secret); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	secret.UserID = user.ID
	if secret.Title == "" || secret.Password == "" {
		h.handleError(w, http.StatusBadRequest, errors.New("title and password cannot be empty"))
		return
	}

	secret.Password, err = crypto.Encrypt(secret.Password, []byte(h.commonConfig.SecretEncKey))
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

	common.RespondWithSuccess(w, http.StatusCreated, "Secret created successfully", map[string]interface{}{})
}

// List handles the retrieval of user secrets.
//
//	@Summary		List user secrets
//	@Description	This endpoint retrieves all secrets associated with the authenticated user.
//	@Tags			secrets
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer <token>"
//	@Success		200				{object}	model.SuccessResponse
//	@Failure		401				{object}	model.ErrorResponse
//	@Failure		500				{object}	model.ErrorResponse
//	@Router			/secret/api/list [get]
func (h *SecretHandler) List(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r, h.commonConfig.JWTSecretKey)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}
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

		decryptedPassword, err := crypto.Decrypt(secrets[i].Password, []byte(h.commonConfig.SecretEncKey))
		if err != nil {
			h.handleError(w, http.StatusInternalServerError, err)
			return
		}
		secrets[i].Password = decryptedPassword
	}

	common.RespondWithSuccess(w, http.StatusOK, "Secrets retrieved successfully", map[string]interface{}{
		"secrets": secrets,
	})
}

// SecretDetail retrieves the details of a specific secret.
//
//	@Summary		Get Secret Details
//	@Description	Retrieves details of a specific secret associated with the authenticated user.
//	@Tags			secrets
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string								true	"Bearer <token>"
//	@Param			secret_id		body		model.SwaggerSecretDetailRequest	true	"ID of the secret to retrieve"
//	@Success		200				{object}	model.SuccessResponse
//	@Failure		400				{object}	model.ErrorResponse
//	@Failure		401				{object}	model.ErrorResponse
//	@Failure		500				{object}	model.ErrorResponse
//	@Router			/secret/api/secretdetail [post]
func (h *SecretHandler) SecretDetail(w http.ResponseWriter, r *http.Request) {
	// Authorization check
	user, err := auth.ValidateToken(r, h.commonConfig.JWTSecretKey)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	var secretDetail model.SecretDetail
	if err := common.ParseAndValidatePayload(r, &secretDetail); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	secret, err := h.service.SecretDetail(r.Context(), user.ID, secretDetail.SecretID)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	common.RespondWithSuccess(w, http.StatusOK, "Secrets detail retrieved successfully", map[string]interface{}{
		"secret": secret,
	})
}

func (h *SecretHandler) Delete(w http.ResponseWriter, r *http.Request) {
	user, err := auth.ValidateToken(r, h.commonConfig.JWTSecretKey)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	idStr := chi.URLParam(r, "id")
	secretID, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteSecretByID(r.Context(), user.ID, secretID)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SecretHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Auth
	user, err := auth.ValidateToken(r, h.commonConfig.JWTSecretKey)
	if err != nil || user == nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	// Parse ID
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Errorf("invalid secret ID"))
		return
	}

	var input model.Secret
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload"))
		return
	}

	err = h.service.UpdateSecret(r.Context(), user.ID, id, input)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	common.RespondWithSuccess(w, http.StatusOK, "Secret updated successfully", map[string]interface{}{
		"secret": input,
	})
}
