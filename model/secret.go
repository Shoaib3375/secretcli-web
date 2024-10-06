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
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// SecretUsecase interface defining methods related to secret use cases
type SecretUsecase interface {
	Create(ctx context.Context, secret Secret) error // Method to create a secret
}

// SecretRepository interface defining methods for interacting with the secret data storage
type SecretRepository interface {
	Create(ctx context.Context, secret Secret) error // Method to save a secret
}
