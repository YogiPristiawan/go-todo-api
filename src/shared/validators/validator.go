package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidatorError is a custom error
// to indicate that tyhe error type is from validator
type ValidatorError struct {
	Err error
}

func (v *ValidatorError) Error() string {
	return fmt.Sprintf("%v", v.Err)
}

func (v *ValidatorError) Unwrap() error {
	return v.Err
}

// CustomErrorMessage make custom validator error message
func CustomErrorMessage(vError error) error {
	if obj, ok := vError.(validator.ValidationErrors); ok {
		for _, vError := range obj {
			switch vError.Tag() {
			case "username":
				return &ValidatorError{
					Err: fmt.Errorf("%s invalid", vError.Field()),
				}
			default:
				return &ValidatorError{
					Err: fmt.Errorf("%v", vError.Error()),
				}
			}
		}
	}

	return vError
}
