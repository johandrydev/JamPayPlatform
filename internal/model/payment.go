// Package model contains the data structures and types used in the JamPay application.
package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "PENDING"
	PaymentStatusSuccess    PaymentStatus = "SUCCESS"
	PaymentStatusFailed     PaymentStatus = "FAILED"
	PaymentStatusRefunded   PaymentStatus = "REFUNDED"
	PaymentStatusRefundFail PaymentStatus = "REFUND_FAIL"
)

type Payment struct {
	ID              uuid.UUID     `json:"id"`
	ExternalID      string        `json:"external_id"`
	MerchantID      uuid.UUID     `json:"merchant_id"`
	CustomerID      uuid.UUID     `json:"customer_id"`
	PaymentMethodID uuid.UUID     `json:"payment_method_id"`
	Amount          float64       `json:"amount"`
	Status          PaymentStatus `json:"status"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	ProcessedAt     *time.Time    `json:"processed_at"`
}

func (p *Payment) IsPending() bool {
	return p.Status == PaymentStatusPending
}

func (p *Payment) IsSuccess() bool {
	return p.Status == PaymentStatusSuccess
}

func changePaymentStatus(p *Payment, status PaymentStatus) {
	now := time.Now()
	if status == PaymentStatusSuccess {
		p.ProcessedAt = &now
	}
	p.Status = status
	p.UpdatedAt = now
}

func (p *Payment) Success() {
	changePaymentStatus(p, PaymentStatusSuccess)
}

func (p *Payment) Fail() {
	changePaymentStatus(p, PaymentStatusFailed)
}

func (p *Payment) Refund() {
	changePaymentStatus(p, PaymentStatusRefunded)
}

func (p *Payment) RefundFail() {
	changePaymentStatus(p, PaymentStatusRefundFail)
}
