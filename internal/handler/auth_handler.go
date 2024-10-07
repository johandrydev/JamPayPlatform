// Package handler provides HTTP handlers for the JamPay application.
package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"JamPay/internal/dto"
	"JamPay/internal/pkg/auth"
	httpJP "JamPay/internal/pkg/http_jp"
	"JamPay/internal/service"
)

// AuthHandler handles authentication-related HTTP requests.
type AuthHandler struct {
	userService *service.UserService
}

// NewAuthHandler creates a new AuthHandler with the given database connection.
func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{
		userService: service.NewUserService(db),
	}
}

// Login handles the user login process.
// It decodes the login input, validates it, finds the user by email,
// checks the password, generates an access token, and writes the response.
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding login input", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, "invalid request body")
	}

	if err := input.Validate(); err != nil {
		log.Println("Error validating login input", err)
		httpJP.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.userService.FindUserByEmail(r.Context(), input.Email)
	if err != nil {
		log.Println("Error finding user by email", err)
		if errors.Is(err, sql.ErrNoRows) {
			httpJP.WriteError(w, r, http.StatusUnauthorized, "invalid email or password")
			return
		}
		httpJP.WriteError(w, r, http.StatusInternalServerError, "error trying to login, please try again later")
		return
	}

	if !user.IsValidPassword(input.Password) {
		httpJP.WriteError(w, r, http.StatusUnauthorized, "invalid email or password")
		return
	}

	token, err := auth.NewAccessToken(user.ID, user.Role)
	if err != nil {
		log.Println("Error generating access token", err)
		httpJP.WriteError(w, r, http.StatusInternalServerError, "error trying to login, please try again later")
		return
	}

	httpJP.WriteJson(w, r, http.StatusOK, dto.LoginOutput{
		Token: token,
	}, "Login successful")
}
