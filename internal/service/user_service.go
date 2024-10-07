// Package service provides the business logic for the JamPay application.
package service

import (
	"context"
	"database/sql"

	"JamPay/internal/model"
	"JamPay/internal/repository"
)

// UserService is a struct that holds a reference to the UserRepository.
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService with the given database connection.
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(db),
	}
}

// FindUserByEmail retrieves a user by their email address.
// It returns a pointer to the User model and an error if any occurs during the query.
func (u *UserService) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.userRepo.FindUserByEmail(ctx, email)
}
