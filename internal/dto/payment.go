package dto

import (
	"fmt"

	"github.com/google/uuid"

	"JamPay/internal/model"
)

type PaymentInput struct {
	Amount          int64     `json:"amount"`
	MerchantID      uuid.UUID `json:"merchant_id"`
	CustomerID      uuid.UUID `json:"customer_id"`
	PaymentMethodID uuid.UUID `json:"payment_method_id"`
}

func (p *PaymentInput) Validate() error {
	if p.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	if p.MerchantID == uuid.Nil {
		return fmt.Errorf("merchant_id is required")
	}
	if p.CustomerID == uuid.Nil {
		return fmt.Errorf("customer_id is required")
	}
	if p.PaymentMethodID == uuid.Nil {
		return fmt.Errorf("payment_method_id is required")
	}
	return nil
}

func (p *PaymentInput) ToPayment() *model.Payment {
	return &model.Payment{
		Amount:          float64(p.Amount),
		MerchantID:      p.MerchantID,
		CustomerID:      p.CustomerID,
		PaymentMethodID: p.PaymentMethodID,
		Status:          model.PaymentStatusPending,
	}
}
