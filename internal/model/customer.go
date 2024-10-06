// Package model contains the data structures and types used in the JamPay application.
package model

import (
	"time"

	"github.com/google/uuid"
)

type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "ACTIVE"
	CustomerStatusInactive CustomerStatus = "INACTIVE"
)

type Customer struct {
	ID             uuid.UUID
	ExternalID     string
	Name           string
	Email          string
	Status         CustomerStatus
	PaymentMethods []PaymentMethod
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
