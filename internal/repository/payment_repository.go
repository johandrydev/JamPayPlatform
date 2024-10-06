// Package repository provides functions to interact with the payment data in the database.
package repository

import (
	"context"
	"database/sql"
	"log"

	"JamPay/internal/model"
)

// PaymentRepository represents a repository for managing payment data.
type PaymentRepository struct {
	db *sql.DB
}

// NewPaymentRepository creates a new instance of PaymentRepository with the provided database connection.
func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// Save inserts a new payment record into the database.
// It returns an error if the insertion fails or if there is an issue with the database query.
func (p *PaymentRepository) Save(ctx context.Context, payment *model.Payment) error {
	// put ID in the payment object
	err := p.db.QueryRowContext(
		ctx, `
			INSERT INTO
			    payments (
			    	external_id,
					merchant_id,
					customer_id,
					payment_method_id,
					amount,
					status
              	)
        	VALUES ($1, $2, $3, $4, $5, $6)
        	RETURNING id, created_at`,
		payment.ExternalID,
		payment.MerchantID,
		payment.CustomerID,
		payment.PaymentMethodID,
		payment.Amount,
		payment.Status,
	).Scan(
		&payment.ID,
		&payment.CreatedAt,
	)
	return err
}

// FindAllByMerchantID retrieves all payments associated with a specific merchant ID from the database.
// It returns a slice of Payment objects and an error if the query fails.
func (p *PaymentRepository) FindAllByMerchantID(ctx context.Context, merchantID string) ([]*model.Payment, error) {
	rows, err := p.db.QueryContext(
		ctx, `
			SELECT
			    id,
			    merchant_id,
			    customer_id,
			    payment_method_id,
			    amount,
			    status,
			    processed_at
			FROM payments WHERE merchant_id = $1`,
		merchantID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*model.Payment
	for rows.Next() {
		payment := &model.Payment{}
		if err := rows.Scan(
			&payment.ID,
			&payment.MerchantID,
			&payment.CustomerID,
			&payment.PaymentMethodID,
			&payment.Amount,
			&payment.Status,
			&payment.ProcessedAt,
		); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// FindByID retrieves a payment by its ID from the database.
// It returns a pointer to the Payment object and an error if the payment is not found or if there is an issue with the database query.
func (p *PaymentRepository) FindByID(ctx context.Context, paymentID string) (*model.Payment, error) {
	payment := &model.Payment{}
	err := p.db.QueryRowContext(
		ctx, `
			SELECT
			    id,
			    merchant_id,
			    customer_id,
			    payment_method_id,
			    external_id,
			    amount,
			    status,
			    processed_at
        FROM payments WHERE id = $1`,
		paymentID,
	).Scan(
		&payment.ID,
		&payment.MerchantID,
		&payment.CustomerID,
		&payment.PaymentMethodID,
		&payment.ExternalID,
		&payment.Amount,
		&payment.Status,
		&payment.ProcessedAt,
	)
	if err != nil {
		log.Println("error finding payment", err)
		return payment, err
	}
	return payment, err
}

// Update modifies an existing payment record in the database.
// It returns an error if the update fails or if there is an issue with the database query.
func (p *PaymentRepository) Update(ctx context.Context, payment *model.Payment) error {
	_, err := p.db.ExecContext(
		ctx, `
			UPDATE payments
			SET
				status = $1,
				updated_at = $2,
				processed_at = $3,
				external_id = $4
			WHERE id = $5`,
		payment.Status,
		payment.UpdatedAt,
		payment.ProcessedAt,
		payment.ExternalID,
		payment.ID,
	)
	return err
}
