package auth

import (
	"context"
	"time"

	"github.com/mahinops/secretcli-web/model"
	"gorm.io/gorm"
)

type SqlAuthRepository struct {
	db *gorm.DB
}

func NewSqlAuthRepository(db *gorm.DB) *SqlAuthRepository {
	return &SqlAuthRepository{db: db}
}

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

func (r *SqlAuthRepository) GetByEmail(ctx context.Context, email string) (*model.Auth, error) {
	var user model.Auth
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SqlAuthRepository) UpdateLastAuth(ctx context.Context, userID uint) error {
	return r.db.Model(&model.Auth{}).Where("id = ?", userID).Update("last_auth", time.Now()).Error
}

func (r *SqlAuthRepository) UpdateExpiry(ctx context.Context, userID uint, jwtExpiryDuration time.Duration) error {
	expiryTime := time.Now().Add(jwtExpiryDuration) // Use the constant here

	return r.db.Model(&model.Auth{}).Where("id = ?", userID).Update("expiry", expiryTime).Error
}
