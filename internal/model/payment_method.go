// Package model contains the data structures and types used in the JamPay application.
package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethodType string

const (
	CreditCard   PaymentMethodType = "CREDIT_CARD"
	DebitCard    PaymentMethodType = "DEBIT_CARD"
	BankTransfer PaymentMethodType = "BANK_TRANSFER"
)

type PaymentMethod struct {
	ID            uuid.UUID
	OwnerID       uuid.UUID
	ExternalID    string
	Type          PaymentMethodType
	ProductNumber string
	ExpiryDate    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// ToStripePaymentMethod converts the PaymentMethod type to a corresponding
// Stripe payment method type. It returns a slice of strings representing
// the Stripe payment method type(s).
func (p *PaymentMethod) ToStripePaymentMethod() []string {
	switch p.Type {
	case CreditCard:
		return []string{"card"}
	case DebitCard:
		return []string{"card"}
	case BankTransfer:
		return []string{"bank_transfer"}
	default:
		return []string{}
	}
}
