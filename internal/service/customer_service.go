package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"JamPay/internal/model"
	"JamPay/internal/repository"
)

type CustomerService struct {
	CustomerRepo *repository.CustomerRepository
}

func NewCustomerService(db *sql.DB) *CustomerService {
	return &CustomerService{CustomerRepo: repository.NewCustomerRepository(db)}
}

func (s *CustomerService) FindByID(ctx context.Context, customerID uuid.UUID) (*model.Customer, error) {
	return s.CustomerRepo.FindByID(ctx, customerID)
}
