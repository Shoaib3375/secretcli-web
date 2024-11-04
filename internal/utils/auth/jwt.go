package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mahinops/secretcli-web/model"
)

func GenerateToken(user *model.Auth, JWTSecretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"exp":       user.Expiry,
		"last_auth": user.LastAuth,
	})

	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(r *http.Request, JWTSecretKey string) (*model.Auth, error) {
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

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.NewValidationError("invalid claims", jwt.ValidationErrorClaimsInvalid)
	}

	expiryStr, ok := claims["exp"].(string)
	if !ok {
		return nil, jwt.NewValidationError("invalid expiry claim", jwt.ValidationErrorClaimsInvalid)
	}

	expiryTime, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()

	if currentTime.After(expiryTime) {
		return nil, jwt.NewValidationError("token has expired", jwt.ValidationErrorExpired)
	}

	user := &model.Auth{
		ID:    uint(claims["user_id"].(float64)), // Type assertion to uint
		Email: claims["email"].(string),
	}
	return user, nil
}

func GetJWTExpiryTime(expiry string) (time.Duration, error) {
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		return 0, err
	}
	return duration, nil
}
