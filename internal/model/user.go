// Package model contains the data structures and types used in the JamPay application.
package model

type UserRoles string

const (
	UserRoleCustomer UserRoles = "CUSTOMER"
	UserRoleMerchant UserRoles = "MERCHANT"
)

type User struct {
	ID             string
	Email          string
	Role           UserRoles
	HashedPassword string
	CreatedAt      string
	UpdatedAt      string
}
