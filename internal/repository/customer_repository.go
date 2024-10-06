// Package repository provides functions to interact with the customer data in the database.
package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	"JamPay/internal/model"
)

// CustomerRepository represents a repository for managing customer data.
type CustomerRepository struct {
	db *sql.DB
}

// NewCustomerRepository creates a new instance of CustomerRepository with the provided database connection.
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// FindByID retrieves a customer by their ID from the database.
// It returns an error if the customer is not found or if there is an issue with the database query.
func (c *CustomerRepository) FindByID(ctx context.Context, customerID uuid.UUID) (*model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRowContext(
		ctx, `
			SELECT
			    id,
			    external_id,
			    name,
			    email,
			    status
			FROM
			    customers
			WHERE
			    id = $1`,
		customerID,
	).Scan(
		&customer.ID,
		&customer.ExternalID,
		&customer.Name,
		&customer.Email,
		&customer.Status,
	)
	if err != nil {
		log.Println("error finding customer", err)
		return &customer, err
	}
	return &customer, err
}

// UpdateStatus updates the status of a customer in the database.
// It returns an error if the update fails or if there is an issue with the database query.
func (c *CustomerRepository) UpdateStatus(customerID string, status string) error {
	// TODO: Implement the logic to update the status of a customer
	return nil
}
