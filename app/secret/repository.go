package secret

import (
	"context"

	"github.com/mahinops/secretcli-web/model"
	"gorm.io/gorm"
)

type SqlSecretRepository struct {
	db *gorm.DB
}

// NewSqlSecretRepository creates a new instance of SqlSecretRepository
func NewSqlSecretRepository(db *gorm.DB) *SqlSecretRepository {
	return &SqlSecretRepository{db: db}
}

// Create method for saving a secret to the database
func (r *SqlSecretRepository) Create(ctx context.Context, secret model.Secret) error {
	// Use the GORM Create method to insert the secret into the database
	return r.db.Create(&secret).Error
}

// List method for retrieving all secrets from the database
func (r *SqlSecretRepository) List(ctx context.Context, userID uint) ([]model.Secret, error) {
	// Use the GORM Find method to retrieve all secrets from the database
	var secrets []model.Secret
	if err := r.db.Where("user_id = ?", userID).Find(&secrets).Error; err != nil {
		return nil, err
	}
	return secrets, nil
}
