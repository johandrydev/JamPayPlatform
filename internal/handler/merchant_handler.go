// Package handler provides HTTP handlers for the JamPay application.
package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"

	httpJP "JamPay/internal/pkg/http_jp"
	"JamPay/internal/service"
)

type MerchantHandler struct {
	MerchantService *service.MerchantService
}

// NewMerchantHandler creates a new MerchantHandler with the given database connection.
func NewMerchantHandler(db *sql.DB) *MerchantHandler {
	return &MerchantHandler{
		MerchantService: service.NewMerchantService(db),
	}
}

// FindMerchant handles the HTTP request to find a merchant by its ID.
// It retrieves the merchant ID from the URL parameters, fetches the merchant
// from the database using the MerchantService, and writes the response as JSON.
func (m *MerchantHandler) FindMerchant(w http.ResponseWriter, r *http.Request) {
	merchantID := chi.URLParam(r, "merchantID")
	merchant, err := m.MerchantService.FindByID(merchantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpJP.WriteJson(w, r, http.StatusOK, merchant, "Merchant information retrieved successfully")
}
