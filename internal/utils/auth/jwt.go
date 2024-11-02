package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mahinops/secretcli-web/model"
)

// Secret key used for signing JWT tokens
var jwtSecret = []byte("your-secret-key")

// GenerateToken generates a JWT token for an authenticated user
func GenerateToken(user *model.Auth) (string, error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"exp":       user.Expiry,
		"last_auth": user.LastAuth,
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token from the request and returns the user information
func ValidateToken(r *http.Request) (*model.Auth, error) {
	// Extract the token from the Authorization header
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return nil, http.ErrNoCookie // Return an error if no token is found
	}

	var tokenString string
	if strings.HasPrefix(bearerToken, "Bearer ") {
		tokenString = strings.Split(bearerToken, " ")[1]
	} else {
		tokenString = bearerToken
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.NewValidationError("invalid claims", jwt.ValidationErrorClaimsInvalid)
	}

	// Validate expiry
	expiryStr, ok := claims["exp"].(string)
	if !ok {
		return nil, jwt.NewValidationError("invalid expiry claim", jwt.ValidationErrorClaimsInvalid)
	}

	// Parse the expiry time string
	expiryTime, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		return nil, err
	}
	// Get the current time
	currentTime := time.Now()

	// Check if the token is still valid
	// Check if the token is still valid
	if currentTime.After(expiryTime) {
		return nil, jwt.NewValidationError("token has expired", jwt.ValidationErrorExpired)
	}

	// Return user information from claims
	user := &model.Auth{
		ID:    uint(claims["user_id"].(float64)), // Type assertion to uint
		Email: claims["email"].(string),
	}
	return user, nil
}
