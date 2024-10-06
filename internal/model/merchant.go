// Package model contains the data structures and types used in the JamPay application.
package model

import "github.com/google/uuid"

type MerchantStatus string

const (
	MerchantStatusPending    MerchantStatus = "PENDING"
	MerchantStatusVerified   MerchantStatus = "VERIFIED"
	MerchantStatusUnverified MerchantStatus = "UNVERIFIED"
)

type Merchant struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	BankAccount string         `json:"bank_account"`
	Status      MerchantStatus `json:"status"`
}
