package validation

import "errors"

func ValidateUpdateUserInput(name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	return nil
}
