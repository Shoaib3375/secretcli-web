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

func (r *SqlAuthRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.Auth{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil // Return true if count is greater than 0
}
