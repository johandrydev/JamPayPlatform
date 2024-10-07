// Package auth provides functions for generating and verifying JWT access tokens.
package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"JamPay/internal/model"
)

// secretKey is the key used to sign the JWT tokens.
var secretKey = []byte("JamPaySecretKey")

// NewAccessToken generates a new JWT access token for the given user ID and role.
// The token expires in 2 hours.
func NewAccessToken(userID uuid.UUID, role model.UserRole) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": map[string]any{
			"user_id": userID,
			"role":    role,
		},
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return claims.SignedString([]byte(secretKey))

}

// VerifyAccessToken verifies the given JWT access token string.
// It returns the parsed token if valid, or an error if invalid.
func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	tokenArray := strings.Split(tokenString, " ")
	if len(tokenArray) != 2 {
		return nil, jwt.ErrSignatureInvalid
	}
	token := tokenArray[1]
	return jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
}
