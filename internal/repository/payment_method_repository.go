// Package repository provides functions to interact with the payment method data in the database.
package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"

	"JamPay/internal/model"
)

// PaymentMethodRepository represents a repository for managing payment method data.
type PaymentMethodRepository struct {
	db *sql.DB
}

// NewPaymentMethodRepository creates a new instance of PaymentMethodRepository with the provided database connection.
func NewPaymentMethodRepository(db *sql.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

// FindById retrieves a payment method by its ID from the database.
// It returns a pointer to the PaymentMethod object and an error if the payment method is not found or if there is an issue with the database query.
func (p *PaymentMethodRepository) FindById(paymentMethodId uuid.UUID) (*model.PaymentMethod, error) {
	paymentMethod := new(model.PaymentMethod)
	err := p.db.QueryRow(`
		SELECT
			id,
			owner_id,
			external_id,
			type,
			product_number,
			expiration_date
		FROM
			payment_methods
		WHERE
			id = $1`,
		paymentMethodId,
	).Scan(
		&paymentMethod.ID,
		&paymentMethod.OwnerID,
		&paymentMethod.ExternalID,
		&paymentMethod.Type,
		&paymentMethod.ProductNumber,
		&paymentMethod.ExpiryDate,
	)
	if err != nil {
		log.Println("error finding payment method", err)
		return nil, err
	}
	return paymentMethod, nil
}
