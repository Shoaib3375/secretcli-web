package model

import (
	"context"
	"time"
)

type Secret struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Note      string     `json:"note"`
	Email     string     `json:"email"`
	Website   string     `json:"website"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type GeneratePasswordRequest struct {
	Length               int  `json:"length"`
	IncludeSpecialSymbol bool `json:"include_special_symbol"`
}

// SecretUsecase interface defining methods related to secret use cases
type SecretUsecase interface {
	Create(ctx context.Context, secret Secret) error
	List(ctx context.Context, userID uint) ([]Secret, error)
	GeneratePassword(ctx context.Context, length int, includeSpecialSymbol bool) (string, error)
}

// SecretRepository interface defining methods for interacting with the secret data storage
type SecretRepository interface {
	Create(ctx context.Context, secret Secret) error
	List(ctx context.Context, userID uint) ([]Secret, error)
	GeneratePassword(ctx context.Context, length int, includeSpecialSymbol bool) (string, error)
}
