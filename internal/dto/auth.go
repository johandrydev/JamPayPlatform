package dto

import "errors"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginInput) Validate() error {
	if l.Email == "" {
		return errors.New("email is required")
	}
	if l.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type LoginOutput struct {
	Token string `json:"token"`
}
