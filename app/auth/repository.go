package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mahinops/secretcli-web/model"
)

// SqlAuthRepository implements the AuthRepository interface
type SqlAuthRepository struct {
	db *sql.DB
}

// NewSqlAuthRepository creates a new SqlAuthRepository
func NewSqlAuthRepository(db *sql.DB) *SqlAuthRepository {
	return &SqlAuthRepository{db: db}
}

// Create saves a new user to the database and returns the user's name and any error encountered
func (r *SqlAuthRepository) Create(ctx context.Context, user model.Auth) (string, error) {

	fmt.Println(user)
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING name`
	var name string
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&name)
	if err != nil {
		return "", err // Return empty string and error if the operation fails
	}
	return name, nil // Return the user's name
}
