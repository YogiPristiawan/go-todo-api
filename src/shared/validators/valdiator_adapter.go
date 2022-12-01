package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// NewValidator creates an instance of valdiator package
func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("username", validateUsername)

	return v
}

// custom validation

// validateUsername is a custom validator
// to validate the given username
func validateUsername(f validator.FieldLevel) bool {
	text := f.Field().String()
	regex, _ := regexp.Compile(`^\w+$`)

	return regex.MatchString(text)
}
