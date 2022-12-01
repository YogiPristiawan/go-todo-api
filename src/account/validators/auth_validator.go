package validators

import (
	"go_todo_api/src/account/dto"

	"github.com/go-playground/validator/v10"
)

// AuthValidator is an abstract that contains methods
// required to handle auth-related validation
type AuthValidator interface {
	ValidateLogin(in dto.LoginRequest) error
	ValidateRegister(in dto.RegisterRequest) error
}

// authValidator provides methods to handle
// auth-related validation
type authValidator struct {
	validator *validator.Validate
}

// NewAuthValidator creates an instance of
// AuthValidator
func NewAuthValidator(validator *validator.Validate) AuthValidator {
	return &authValidator{
		validator: validator,
	}
}

func (a *authValidator) ValidateLogin(in dto.LoginRequest) error {
	return customErrorMsg(a.validator.Struct(&in))
}

func (a *authValidator) ValidateRegister(in dto.RegisterRequest) error {
	return customErrorMsg(a.validator.Struct(in))
}
