package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	merchant "JamPay/internal/handler"
	"JamPay/internal/payment_services/stripe"
	"JamPay/internal/pkg/database"
)

const port = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.CreatePostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	merchanHandlers := merchant.NewMerchantHandler(db)

	stripeProvider := stripe.NewStripeService()
	paymentHandlers := merchant.NewPaymentHandler(db, stripeProvider)

	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Route("/merchant", func(r chi.Router) {
			r.Get("/{merchantID}", merchanHandlers.FindMerchant)
			r.Get("/{merchantID}/payments", paymentHandlers.GetAllByMerchantID)
		})

		r.Route("/payment", func(r chi.Router) {
			r.Post("/", paymentHandlers.CreatePayment)
			r.Get("/{paymentID}", paymentHandlers.GetPayment)
			r.Post("/{paymentID}/process", paymentHandlers.ProcessPayment)
			r.Post("/{paymentID}/refund", paymentHandlers.RefundPayment)
		})
	})

	log.Println("Server is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
