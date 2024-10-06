// Package stripe provides functionality to interact with the Stripe API for payment processing.
package stripe

import (
	"os"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/client"
)

// Service represents a Stripe service client.
type Service struct {
	client *client.API
}

// NewStripeService initializes a new Stripe service client with the provided API key.
func NewStripeService() *Service {
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		panic("STRIPE_SECRET_KEY environment variable not set")
	}

	return &Service{
		client: client.New(key, nil),
	}
}

// CreatePaymentIntent creates a payment intent on Stripe with the specified amount, payment methods, and payment ID.
// It returns the created PaymentIntent or an error if the creation fails.
func (s *Service) CreatePaymentIntent(amount int64, paymentMethod []string, paymentID string, customerID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice(paymentMethod),
		PaymentMethod:      stripe.String(paymentID),
		Customer:           stripe.String(customerID),
		Confirm:            stripe.Bool(true),
		CaptureMethod:      stripe.String(string(stripe.PaymentIntentCaptureMethodAutomatic)),
	}

	return s.client.PaymentIntents.New(params)
}

// RefundPaymentIntent creates a refund for the specified payment intent ID on Stripe.
// It returns the created Refund or an error if the creation fails.
func (s *Service) RefundPaymentIntent(paymentIntentID string) (*stripe.Refund, error) {
	return s.client.Refunds.New(&stripe.RefundParams{
		PaymentIntent: stripe.String(paymentIntentID),
	})
}
