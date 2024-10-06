package service

import (
	"database/sql"

	"JamPay/internal/repository"
)

type PaymentMethodService struct {
	PaymentMethodRepo *repository.PaymentMethodRepository
}

func NewPaymentMethodService(db *sql.DB) *PaymentMethodService {
	return &PaymentMethodService{
		PaymentMethodRepo: repository.NewPaymentMethodRepository(db),
	}
}
