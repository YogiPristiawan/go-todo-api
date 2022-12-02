package validators

import (
	"database/sql/driver"
	"go_todo_api/src/shared/entities"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// NewValidator creates an instance of valdiator package
func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("username", validateUsername)
	v.RegisterCustomTypeFunc(validateValuer, entities.String{})

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

func validateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}
	return nil
}
