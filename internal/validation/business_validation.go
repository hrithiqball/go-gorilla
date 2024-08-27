package validation

import (
	"errors"
)

type CreateBusinessForm struct {
	Name         string `schema:"name" validate:"required"`
	Email        string `schema:"email" validate:"required"`
	Phone        string `schema:"phone"`
	Address      string `schema:"address"`
	Website      string `schema:"website"`
	CoverPhoto   string `schema:"coverPhoto"`
	ProfilePhoto string `schema:"profilePhoto"`
}

func ValidateCreateBusinessFormInput(name, email, phone, address, website string) error {
	if name == "" {
		return errors.New("name is required")
	}

	if email == "" {
		return errors.New("email is required")
	}

	if phone == "" {
		return errors.New("phone is required")
	}

	return nil
}
