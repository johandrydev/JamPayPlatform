// Package middleware provides HTTP middleware for the JamPay application.
package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"JamPay/internal/model"
	"JamPay/internal/pkg/auth"
	httpJP "JamPay/internal/pkg/http_jp"
)

// IsAuthenticated is a middleware that checks if the request has a valid JWT token.
// If the token is valid, it adds the token claims to the request context.
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			httpJP.WriteError(w, r, http.StatusUnauthorized, "missing token")
			return
		}

		claims, err := auth.VerifyAccessToken(token)
		if err != nil {
			httpJP.WriteError(w, r, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "jam_pay_claims", claims.Claims.(jwt.MapClaims))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// IsMerchant is a middleware that checks if the user has a merchant role.
// If the user is not a merchant, it returns a 401 Unauthorized response.
func IsMerchant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the claims from the context
		claims := r.Context().Value("jam_pay_claims").(jwt.MapClaims)
		if claims["sub"].(map[string]interface{})["role"] != string(model.UserRoleMerchant) {
			httpJP.WriteError(w, r, http.StatusUnauthorized, "user is not authorized to perform this action")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsCustomer is a middleware that checks if the user has a customer role.
// If the user is not a customer, it returns a 401 Unauthorized response.
func IsCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the claims from the context
		claims := r.Context().Value("jam_pay_claims").(jwt.MapClaims)
		if claims["sub"].(map[string]interface{})["role"] != string(model.UserRoleCustomer) {
			httpJP.WriteError(w, r, http.StatusUnauthorized, "user is not authorized to perform this action")
			return
		}

		next.ServeHTTP(w, r)
	})
}
