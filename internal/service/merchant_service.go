package service

import (
	"database/sql"

	"JamPay/internal/model"
	"JamPay/internal/repository"
)

type MerchantService struct {
	MerchantRepo *repository.MerchantRepository
}

func NewMerchantService(db *sql.DB) *MerchantService {
	return &MerchantService{MerchantRepo: repository.NewMerchantRepository(db)}
}

func (s *MerchantService) FindByID(merchantID string) (*model.Merchant, error) {
	return s.MerchantRepo.FindByID(merchantID)
}
