package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Validate is an abstract that contains methods
// to handle validations
type Validate interface {
	Struct(interface{}) error
}

// CustomErrorMessage make custom validator error message
func CustomErrorMessage(vError error) error {
	if obj, ok := vError.(validator.ValidationErrors); ok {
		for _, vError := range obj {
			switch vError.Tag() {
			case "username":
				return fmt.Errorf("%s invalid", vError.Field())
			}
		}
	}
	return nil
}
