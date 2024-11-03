package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/auth"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/model"
)

type AuthHandler struct {
	usecase     model.AuthUsecase
	renderer    *tmplrndr.Renderer
	redisClient *redis.Client
}

func NewAuthHandler(usecase model.AuthUsecase, renderer *tmplrndr.Renderer, redisClient *redis.Client) *AuthHandler {
	return &AuthHandler{usecase: usecase, renderer: renderer, redisClient: redisClient}
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
func (h *AuthHandler) handleError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Code:    code,
		Message: err.Error(),
	})
}

// RegisterUser handles user registration.
//
//	@Summary		Register a new user
//	@Description	This endpoint allows a new user to register with a name, email and password.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.SwaggerAuthRequest	true	"User registration payload"
//	@Success		201		{object}	model.SuccessResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		409		{object}	model.ErrorResponse
//	@Router			/auth/api/register [post]
func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.Auth
	if err := common.ParseAndValidatePayload(r, &user); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	name, err := h.usecase.Create(r.Context(), user)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "email already exists" {
			statusCode = http.StatusConflict
		}
		h.handleError(w, statusCode, err)
		return
	}

	common.RespondWithSuccess(w, http.StatusCreated, "User created successfully", map[string]interface{}{
		"name": name,
	})
}

// LoginUser handles user login.
//
//	@Summary		Login a user
//	@Description	This endpoint allows a user to login with an email and password.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.SwaggerUserLoginRequest	true	"User login payload"
//	@Success		200		{object}	model.SuccessResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		401		{object}	model.ErrorResponse
//	@Router			/auth/api/login [post] // Update this to the correct login endpoint
func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest model.UserLogin
	if err := common.ParseAndValidatePayload(r, &loginRequest); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.usecase.Login(r.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	// Generate a JWT token for the user
	token, err := auth.GenerateToken(user)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	h.redisClient.Set(r.Context(), token, user.Email, 1*time.Minute)

	common.RespondWithSuccess(w, http.StatusOK, "Login successful", map[string]interface{}{
		"token":  token,
		"expiry": user.Expiry,
	})
}
