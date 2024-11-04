package secret

import (
	"context"
	"errors"

	"github.com/mahinops/secretcli-web/model"
)

type SecretService struct {
	repo model.SecretRepository
}

func NewSecretService(repo model.SecretRepository) *SecretService {
	return &SecretService{repo: repo}
}

func (s *SecretService) Create(ctx context.Context, secret model.Secret) error {
	if secret.Title == "" || secret.Password == "" {
		return errors.New("title and password cannot be empty")
	}
	return s.repo.Create(ctx, secret)
}

func (s *SecretService) List(ctx context.Context, userID uint) ([]model.Secret, error) {
	return s.repo.List(ctx, userID)
}

func (s *SecretService) GeneratePassword(ctx context.Context, length int, includeSpecialSymbol bool) (string, error) {
	return s.repo.GeneratePassword(ctx, length, includeSpecialSymbol)
}

func (s *SecretService) SecretDetail(ctx context.Context, userID uint, SecretID int) (model.Secret, error) {
	if userID < 1 || SecretID < 1 {
		return model.Secret{}, errors.New("user id and secret id cannot be negative")
	}
	return s.repo.SecretDetail(ctx, userID, SecretID)
}
