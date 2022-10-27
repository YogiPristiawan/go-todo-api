package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// valdiatorAdapter is a adapter pattern struct
// to handle action about validation with
// golang validator library
type validatorAdapter struct {
	validator *validator.Validate
}

// NewValidatorAdapter creates an instance of valdiatorAdapter struct
func NewValidatorAdapter() Validate {
	v := validator.New()
	v.RegisterValidation("username", validateUsername)

	return &validatorAdapter{}
}

// Struct validate struct object type
func (v *validatorAdapter) Struct(s interface{}) (err error) {
	err = v.validator.Struct(&s)
	return
}

// custom validation

// validateUsername is a custom validator
// to validate the given username
func validateUsername(f validator.FieldLevel) bool {
	text := f.Field().String()
	regex, _ := regexp.Compile(`^\w+$`)

	return regex.MatchString(text)
}
