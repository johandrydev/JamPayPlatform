package service

import (
	"context"

	"JamPay/internal/repository"
)

type CustomerService struct {
	CustomerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) *CustomerService {
	return &CustomerService{CustomerRepo: customerRepo}
}

func (s *CustomerService) FindByID(ctx context.Context, customerID string) error {
	return s.CustomerRepo.FindByID(ctx, customerID)
}
