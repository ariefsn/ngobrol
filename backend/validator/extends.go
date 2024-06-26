package validator

import (
	"regexp"

	vl "github.com/go-playground/validator/v10"
)

func ValidatePassword(fl vl.FieldLevel) bool {
	pattern := `^[a-zA-Z0-9!@#$%^&.?/=-_]*$`
	regex := regexp.MustCompile(pattern)

	val := fl.Field().String()

	if len(val) < 8 {
		return false
	}

	if !regex.MatchString(val) {
		return false
	}

	return true
}
