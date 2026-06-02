package validator

import (
	"errors"
	"regexp"
)

var (
	emailRegex     = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	digitRegex     = regexp.MustCompile(`[0-9]`)
)

type AuthValidator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
}

type authValidator struct{}

func NewAuthValidator() AuthValidator {
	return &authValidator{}
}

func (v *authValidator) ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func (v *authValidator) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if !uppercaseRegex.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !digitRegex.MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	return nil
}
