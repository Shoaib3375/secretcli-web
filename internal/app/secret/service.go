package secret

import (
	"context"
	"errors"

	"github.com/mahinops/secretcli-web/model"
)

type SecretService struct {
	repo model.SecretRepository
}

// NewSecretService creates a new instance of SecretService
func NewSecretService(repo model.SecretRepository) *SecretService {
	return &SecretService{repo: repo}
}

// Create method for secret service
func (s *SecretService) Create(ctx context.Context, secret model.Secret) error {
	// Validate input
	if secret.Title == "" || secret.Password == "" {
		return errors.New("title and password cannot be empty")
	}

	// Call the repository to save the secret
	return s.repo.Create(ctx, secret)
}

// List method for secret service
func (s *SecretService) List(ctx context.Context, userID uint) ([]model.Secret, error) {
	// Call the repository to retrieve all secrets
	return s.repo.List(ctx, userID)
}

func (s *SecretService) GeneratePassword(ctx context.Context, length int, includeSpecialSymbol bool) (string, error) {
	// Call the repository to generate a password
	return s.repo.GeneratePassword(ctx, length, includeSpecialSymbol)
}
