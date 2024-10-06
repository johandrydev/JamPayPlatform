// Package service provides the business logic for handling payments.
package service

import (
	"context"
	"database/sql"
	"log"

	"JamPay/internal/model"
	"JamPay/internal/payment_services/stripe"
	"JamPay/internal/repository"
)

// PaymentService handles payment-related operations.
type PaymentService struct {
	paymentRepo       *repository.PaymentRepository
	paymentMethodRepo *repository.PaymentMethodRepository
	stripeClient      *stripe.Service
}

// NewPaymentService creates a new instance of PaymentService with the provided database connection and Stripe service.
func NewPaymentService(db *sql.DB, service *stripe.Service) *PaymentService {
	return &PaymentService{
		paymentRepo:       repository.NewPaymentRepository(db),
		paymentMethodRepo: repository.NewPaymentMethodRepository(db),
		stripeClient:      service,
	}
}

// Save stores a new payment record in the database.
func (p *PaymentService) Save(ctx context.Context, payment *model.Payment) error {
	return p.paymentRepo.Save(ctx, payment)
}

// FindAllByMerchantID retrieves all payments associated with a specific merchant ID.
func (p *PaymentService) FindAllByMerchantID(ctx context.Context, merchantID string) ([]*model.Payment, error) {
	return p.paymentRepo.FindAllByMerchantID(ctx, merchantID)
}

// FindByID retrieves a payment by its ID.
func (p *PaymentService) FindByID(ctx context.Context, id string) (*model.Payment, error) {
	return p.paymentRepo.FindByID(ctx, id)
}

// Update modifies an existing payment record in the database.
func (p *PaymentService) Update(ctx context.Context, payment *model.Payment) error {
	return p.paymentRepo.Update(ctx, payment)
}

// ProcessPayment processes a payment using the Stripe service.
func (p *PaymentService) ProcessPayment(ctx context.Context, payment *model.Payment, customer *model.Customer) error {
	amountInCents := int64(payment.Amount * 100)
	paymentMethod, err := p.paymentMethodRepo.FindById(payment.PaymentMethodID)
	if err != nil {
		log.Println("error finding payment method", err)
		return err
	}
	stripePaymentMethod := paymentMethod.ToStripePaymentMethod()
	providerResult, err := p.stripeClient.CreatePaymentIntent(amountInCents, stripePaymentMethod, paymentMethod.ExternalID, customer.ExternalID)
	if err != nil {
		log.Println("error creating payment intent", err)
		payment.Fail()
		return p.Update(ctx, payment)
	}
	payment.ExternalID = providerResult.ID
	payment.Success()
	return p.Update(ctx, payment)
}

// RefundPayment refunds a payment using the Stripe service.
func (p *PaymentService) RefundPayment(ctx context.Context, payment *model.Payment) error {
	refund, err := p.stripeClient.RefundPaymentIntent(payment.ExternalID)
	if err != nil {
		log.Println("error refunding payment intent", err)
		payment.RefundFail()
		return p.Update(ctx, payment)
	}
	if refund.Status != "succeeded" {
		log.Println("refund failed", refund.FailureReason)
		payment.RefundFail()
		return p.Update(ctx, payment)
	}
	payment.Refund()
	return p.Update(ctx, payment)
}
