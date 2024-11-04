package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/mahinops/secretcli-web/internal/utils/auth"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/model"
)

type AuthHandler struct {
	usecase      model.AuthUsecase
	redisClient  *redis.Client
	commonConfig *common.CommonConfig
}

func NewAuthHandler(usecase model.AuthUsecase, redisClient *redis.Client, commonConfig *common.CommonConfig) *AuthHandler {
	return &AuthHandler{usecase: usecase, redisClient: redisClient, commonConfig: commonConfig}
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

	jwtExpiryDuration, err := auth.GetJWTExpiryTime(h.commonConfig.JWTExpiry)
	if err != nil {
		fmt.Println("Error calculating JWT expiry time:", err)
		return
	}

	user, err := h.usecase.Login(r.Context(), loginRequest.Email, loginRequest.Password, jwtExpiryDuration)
	if err != nil {
		h.handleError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(user, h.commonConfig.JWTSecretKey)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.redisClient.Set(r.Context(), token, user.Email, jwtExpiryDuration).Err()
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
	}
	common.RespondWithSuccess(w, http.StatusOK, "Login successful", map[string]interface{}{
		"token":  token,
		"expiry": user.Expiry,
	})
}
