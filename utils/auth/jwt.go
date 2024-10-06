package auth

import (
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
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
