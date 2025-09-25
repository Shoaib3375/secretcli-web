package secret

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"

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

func (s *SecretService) DeleteSecret(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	userID := r.Context().Value("user_id").(uint)

	err := s.repo.DeleteSecretByID(r.Context(), userID, id)
	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (r *SqlSecretRepository) UpdateSecret(ctx context.Context, userID uint, id int, input model.Secret) error {
	result := r.db.Model(&model.Secret{}).
		Where("user_id = ? AND id = ?", userID, id).
		Updates(map[string]interface{}{
			"title":    input.Title,
			"username": input.Username,
			"password": input.Password,
			"url":      input.URL,
			"note":     input.Note,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("secret not found or not updated")
	}

	return nil
}
func (s *SecretService) DeleteSecretByID(ctx context.Context, userID uint, secretID int) error {
	return s.repo.DeleteSecretByID(ctx, userID, secretID)
}
