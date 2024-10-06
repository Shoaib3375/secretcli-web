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
