// Package repository provides functions to interact with the merchant data in the database.
package repository

import (
	"database/sql"

	"JamPay/internal/model"
)

// MerchantRepository represents a repository for managing merchant data.
type MerchantRepository struct {
	db *sql.DB
}

// NewMerchantRepository creates a new instance of MerchantRepository with the provided database connection.
func NewMerchantRepository(db *sql.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

// FindByID retrieves a merchant by their ID from the database.
// It returns a pointer to the Merchant object and an error if the merchant is not found or if there is an issue with the database query.
func (m *MerchantRepository) FindByID(merchantID string) (*model.Merchant, error) {
	merchant := new(model.Merchant)
	err := m.db.QueryRow(`
		SELECT
    		id,
    		name,
    		email,
    		bank_account,
    		status
		FROM 
	    	merchants
		WHERE 
	    id = $1`,
		merchantID,
	).Scan(
		&merchant.ID,
		&merchant.Name,
		&merchant.Email,
		&merchant.BankAccount,
		&merchant.Status,
	)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}
