package secret

import (
	"context"
	"fmt"
	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/model"
	"gorm.io/gorm"
)

type SqlSecretRepository struct {
	db *gorm.DB
}

func NewSqlSecretRepository(db *gorm.DB) *SqlSecretRepository {
	return &SqlSecretRepository{db: db}
}

func (r *SqlSecretRepository) Create(ctx context.Context, secret model.Secret) error {
	return r.db.Create(&secret).Error
}

func (r *SqlSecretRepository) List(ctx context.Context, userID uint) ([]model.Secret, error) {
	var secrets []model.Secret
	if err := r.db.Where("user_id = ?", userID).Find(&secrets).Error; err != nil {
		return nil, err
	}
	return secrets, nil
}

func (r *SqlSecretRepository) GeneratePassword(ctx context.Context, length int, includeSpecialSymbol bool) (string, error) {
	passwordGenerated, err := crypto.GeneratePassword(length, includeSpecialSymbol)
	if err != nil {
		return "", err
	}
	return passwordGenerated, nil
}

func (r *SqlSecretRepository) SecretDetail(ctx context.Context, userID uint, secretID int) (model.Secret, error) {
	var secret model.Secret
	if err := r.db.Where("user_id = ? AND id = ?", userID, secretID).First(&secret).Error; err != nil {
		return model.Secret{}, err
	}
	return secret, nil
}

func (r *SqlSecretRepository) DeleteSecretByID(ctx context.Context, userID uint, secretID int) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, secretID).Delete(&model.Secret{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("secret not found or not deleted")
	}
	return nil
}

func (s *SecretService) UpdateSecret(ctx context.Context, userID uint, id int, input model.Secret) error {
	return s.repo.UpdateSecret(ctx, userID, id, input)
}
