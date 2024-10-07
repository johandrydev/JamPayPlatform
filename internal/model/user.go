// Package model contains the data structures and types used in the JamPay application.
package model

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	UserRoleCustomer UserRole = "CUSTOMER"
	UserRoleMerchant UserRole = "MERCHANT"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	Role           UserRole  `json:"role"`
	HashedPassword string    `json:"-"`
	CreatedAt      string    `json:"createdAt"`
	UpdatedAt      string    `json:"updatedAt"`
}

func (u *User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	if err != nil {
		log.Println("Error comparing password:", err)
		return false
	}
	return true
}
