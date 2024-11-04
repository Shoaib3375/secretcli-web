package auth

import (
	"context"
	"errors"
	"time"

	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/model"
)

type AuthService struct {
	repo model.AuthRepository
}

func NewAuthService(repo model.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(ctx context.Context, user model.Auth) (string, error) {
	if user.Email == "" || user.Password == "" {
		return "", errors.New("email and password cannot be empty")
	}

	exists, err := s.repo.EmailExists(ctx, user.Email)
	if err != nil {
		return "", err // Handle error accordingly
	}
	if exists {
		return "", errors.New("email already exists") // Return an appropriate error
	}

	hashedPassword, err := crypto.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = &user.CreatedAt
	return s.repo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string, jwtExpiryDuration time.Duration) (*model.Auth, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := crypto.VerifyPassword(user.Password, password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := s.repo.UpdateLastAuth(ctx, user.ID); err != nil {
		return nil, errors.New("failed to update last authentication time")
	}
	if err := s.repo.UpdateExpiry(ctx, user.ID, jwtExpiryDuration); err != nil {
		return nil, errors.New("failed to update expiry time")
	}
	return user, nil
}
