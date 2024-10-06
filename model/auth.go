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
}

type AuthRepository interface {
	Create(ctx context.Context, user Auth) (string, error)
}
