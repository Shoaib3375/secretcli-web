package auth

import (
	"context"
	"errors"
	"time"

	"github.com/mahinops/secretcli-web/model"
	"github.com/mahinops/secretcli-web/utils/crypto"
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

func (s *AuthService) Login(ctx context.Context, email, password string) (*model.Auth, error) {
	// Fetch user by email
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare the provided password with the stored hashed password
	if err := crypto.VerifyPassword(user.Password, password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Update last login time and set new expiry
	user.LastAuth = time.Now()
	user.Expiry = time.Now().Add(24 * time.Hour)

	// Update the last authentication time and expiry in the database
	if err := s.repo.UpdateLastAuth(ctx, user.ID, user.LastAuth); err != nil {
		return nil, errors.New("failed to update last authentication time")
	}
	if err := s.repo.UpdateExpiry(ctx, user.ID, user.Expiry); err != nil {
		return nil, errors.New("failed to update expiry time")
	}
	return user, nil
}
