package model

import (
	"context"
	"time"
)

type Auth struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	LastAuth  time.Time  `json:"last_auth"`
	Expiry    time.Time  `json:"expiry"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type AuthUsecase interface {
	Create(ctx context.Context, user Auth) (string, error)
	Login(ctx context.Context, email, password string) (*Auth, error)
}

type AuthRepository interface {
	Create(ctx context.Context, user Auth) (string, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*Auth, error)
	UpdateLastAuth(ctx context.Context, userID uint, lastAuth time.Time) error
	UpdateExpiry(ctx context.Context, id uint, expiry time.Time) error
}
