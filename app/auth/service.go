package auth

import (
	"context"
	"errors"
	"time"

	"github.com/mahinops/secretcli-web/crypto"
	"github.com/mahinops/secretcli-web/model"
)

type AuthService struct {
	repo model.AuthRepository
}

func NewAuthService(repo model.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(ctx context.Context, user model.Auth) (string, error) {
	// Validate input
	if user.Email == "" || user.Password == "" {
		return "", errors.New("email and password cannot be empty")
	}

	// Check if email already exists
	exists, err := s.repo.EmailExists(ctx, user.Email)
	if err != nil {
		return "", err // Handle error accordingly
	}
	if exists {
		return "", errors.New("email already exists") // Return an appropriate error
	}

	// Hash password
	hashedPassword, err := crypto.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = &user.CreatedAt

	return s.repo.Create(ctx, user)
}
