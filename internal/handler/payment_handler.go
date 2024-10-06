// Package handler provides HTTP handlers for the JamPay application.
package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"JamPay/internal/dto"
	"JamPay/internal/payment_services/stripe"
	httpJP "JamPay/internal/pkg/http_jp"
	"JamPay/internal/service"
)

// PaymentHandler handles HTTP requests related to payments.
type PaymentHandler struct {
	paymentService *service.PaymentService
}

// NewPaymentHandler creates a new PaymentHandler with the given database connection and payment provider (Stripe).
func NewPaymentHandler(db *sql.DB, provider *stripe.Service) *PaymentHandler {
	return &PaymentHandler{
		paymentService: service.NewPaymentService(db, provider),
	}
}

// GetAllByMerchantID handles the HTTP request to retrieve all payments associated with a specific merchant ID.
// It extracts the merchant ID from the URL parameters, fetches the payments from the database using the PaymentService,
// and writes the response as JSON.
func (p *PaymentHandler) GetAllByMerchantID(w http.ResponseWriter, r *http.Request) {
	merchantID := chi.URLParam(r, "merchant_id")
	payments, err := p.paymentService.FindAllByMerchantID(r.Context(), merchantID)
	if err != nil {
		log.Println("Error trying to find payments", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, "internal server error")
		return
	}
	httpJP.WriteJson(w, r, http.StatusOK, payments, "Payments retrieved successfully")
}

// GetPayment handles the HTTP request to retrieve a payment by its ID.
// It extracts the payment ID from the URL parameters, fetches the payment from the database using the PaymentService,
// and writes the response as JSON.
func (p *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := chi.URLParam(r, "payment_id")
	payment, err := p.paymentService.FindByID(r.Context(), paymentID)
	if err != nil {
		log.Println("Error trying to find payment", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, "payment not found")
		return
	}
	httpJP.WriteJson(w, r, http.StatusOK, payment, "Payment retrieved successfully")
}

// CreatePayment handles the HTTP request to create a new payment.
// It decodes the request body into a PaymentInput, validates it, converts it to a Payment model,
// saves it using the PaymentService, and writes the response as JSON.
func (p *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var input dto.PaymentInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error trying to decode request body", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := input.Validate(); err != nil {
		httpJP.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	payment := input.ToPayment()
	if err := p.paymentService.Save(r.Context(), payment); err != nil {
		log.Println("Error trying to save payment", err)
		httpJP.WriteError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	httpJP.WriteJson(w, r, http.StatusCreated, payment, "Payment processed successfully")
}

// ProcessPayment handles the HTTP request to process an existing payment.
// It retrieves the payment ID from the URL parameters, fetches the payment from the database,
// processes the payment using the PaymentService, and writes the response as JSON.
func (p *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := chi.URLParam(r, "payment_id")
	payment, err := p.paymentService.FindByID(r.Context(), paymentID)
	if err != nil {
		log.Println("Error trying to find payment", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, "payment not found")
		return
	}
	if !payment.IsPending() {
		httpJP.WriteError(w, r, http.StatusBadRequest, "payment has already been processed")
		return
	}
	err = p.paymentService.ProcessPayment(r.Context(), payment)
	if err != nil {
		log.Println("Error trying to process payment", err)
		httpJP.WriteError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	httpJP.WriteJson(w, r, http.StatusOK, payment, "Payment processed successfully")
}

// RefundPayment handles the HTTP request to refund an existing payment.
// It retrieves the payment ID from the URL parameters, fetches the payment from the database,
// refunds the payment using the PaymentService, and writes the response as JSON.
func (p *PaymentHandler) RefundPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := chi.URLParam(r, "payment_id")
	payment, err := p.paymentService.FindByID(r.Context(), paymentID)
	if err != nil {
		log.Println("Error trying to find payment", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, "payment not found")
		return
	}
	if !payment.IsSuccess() {
		httpJP.WriteError(w, r, http.StatusBadRequest, "payment cannot be refunded")
		return
	}
	err = p.paymentService.RefundPayment(r.Context(), payment)
	if err != nil {
		log.Println("Error trying to refund payment", err)
		httpJP.WriteError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	httpJP.WriteJson(w, r, http.StatusOK, payment, "Payment refunded successfully")
}
