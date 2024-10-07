package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	merchant "JamPay/internal/handler"
	"JamPay/internal/payment_services/stripe"
	"JamPay/internal/pkg/database"
	"JamPay/internal/pkg/middleware"
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

	authHandler := merchant.NewAuthHandler(db)

	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Post("/login", authHandler.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.IsAuthenticated)

			r.Route("/merchant", func(r chi.Router) {
				r.Get("/{merchantID}", merchanHandlers.FindMerchant)

				r.With(middleware.IsMerchant).Route("/{merchantID}/payments", func(r chi.Router) {
					r.Get("/", paymentHandlers.GetAllByMerchantID)
				})
			})

			r.Route("/payment", func(r chi.Router) {
				r.With(middleware.IsCustomer).Post("/", paymentHandlers.CreatePayment)
				r.Route("/{paymentID}", func(r chi.Router) {
					r.Get("/", paymentHandlers.GetPayment)

					r.With(middleware.IsMerchant).Post("/process", paymentHandlers.ProcessPayment)
					r.With(middleware.IsMerchant).Post("/refund", paymentHandlers.RefundPayment)
				})
			})
		})
	})

	log.Println("Server is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
