package auth

import (
	"context"

	"github.com/mahinops/secretcli-web/model"
	"gorm.io/gorm"
)

// SqlAuthRepository implements the AuthRepository interface
type SqlAuthRepository struct {
	db *gorm.DB
}

// NewSqlAuthRepository creates a new SqlAuthRepository
func NewSqlAuthRepository(db *gorm.DB) *SqlAuthRepository {
	return &SqlAuthRepository{db: db}
}

// Create saves a new user to the database and returns the user's name and any error encountered// Create implements the AuthRepository interface
func (r *SqlAuthRepository) Create(ctx context.Context, user model.Auth) (string, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return "", err
	}
	return user.Name, nil
}
