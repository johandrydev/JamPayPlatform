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

	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Route("/merchant", func(r chi.Router) {
			merchanHandlers := merchant.NewMerchantHandler(db)
			r.Get("/{merchantID}", merchanHandlers.FindMerchant)
		})

		r.Route("/payment", func(r chi.Router) {
			stripeProvider := stripe.NewStripeService()
			paymentHandlers := merchant.NewPaymentHandler(db, stripeProvider)
			r.Get("/merchant/{merchant_id}", paymentHandlers.GetAllByMerchantID)
			r.Get("/{payment_id}", paymentHandlers.GetPayment)
			r.Post("/", paymentHandlers.CreatePayment)
			r.Post("/process/{payment_id}", paymentHandlers.ProcessPayment)
			r.Post("/refund/{payment_id}", paymentHandlers.RefundPayment)
		})
	})

	log.Println("Server is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
