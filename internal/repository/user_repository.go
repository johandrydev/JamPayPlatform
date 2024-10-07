// Package repository provides functions to interact with the user data in the database.
package repository

import (
	"context"
	"database/sql"

	"JamPay/internal/model"
)

// UserRepository is a struct that holds a reference to the database connection.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository with the given database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// FindUserByEmail retrieves a user from the database by their email address.
// It returns a pointer to the User model and an error if any occurs during the query.
func (u *UserRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRowContext(
		ctx,
		`SELECT
            id,
            email,
            role,
            hashed_password
		FROM users
		WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.HashedPassword,
	)
	return &user, err
}
