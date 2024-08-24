package validation

import (
	"regexp"
	"strconv"
)

func IsValidId(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
