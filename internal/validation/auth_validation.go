package validation

import (
	"errors"
)

func ValidateRegisterInput(email, password, name string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(email) {
		return errors.New("invalid email")
	}

	if password == "" {
		return errors.New("password is required")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if name == "" {
		return errors.New("name is required")
	}

	return nil
}
