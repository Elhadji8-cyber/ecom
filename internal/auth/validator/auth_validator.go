package validator

import (
	"errors"
	"regexp"
)

type AuthValidator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
}

type authValidator struct{}

func NewAuthValidator() AuthValidator {
	return &authValidator{}
}

func (b *authValidator) ValidateEmail(email string) error {
	// Simple regex for email validation
	emailRegex := regexp.MustCompile(`^[a-w0-9._%+\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func (b *authValidator) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	//Check for at lease one uppercase letter
	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	//check for at lease one number
	digit := regexp.MustCompile(`[0-9]`)
	if !digit.MatchString(password) {
		return errors.New("password must contain at least one digit")
	}
	return nil
}
