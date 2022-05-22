package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CreateRequestValidator() *validator.Validate {
	// create validator instance
	validator := validator.New()
	validator.RegisterValidation("username", validateUsername)

	return validator
}

// custom validation
func validateUsername(f validator.FieldLevel) bool {
	text := f.Field().String()
	regex, _ := regexp.Compile(`^\w+$`)

	return regex.MatchString(text)
}
