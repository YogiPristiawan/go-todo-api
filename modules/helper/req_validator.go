package helper

import (
	"fmt"
	"reflect"

	"go_todo_api/modules/exceptions"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func ValidatorErrorTranslate(err error, t ut.Translator) error {
	if he, ok := err.(validator.ValidationErrors); ok {
		fmt.Println(he)
		errors := he.Translate(t)

		for _, val := range errors {
			return exceptions.NewInvariantError(val)
		}
	}
	return err
}

func ValidateRequestIfExists(
	validator *validator.Validate,
	s reflect.Type,
	tag string,
	fields map[string]interface{},
) error {
	for i := 0; i < s.NumField(); i++ {
		sField := s.Field(i)
		if alias, ok := sField.Tag.Lookup(tag); ok {
			if alias != "" {
				if val, ok := fields[alias]; ok {

					if validate, ok := sField.Tag.Lookup("validate"); ok {
						if err := validator.Var(val, validate); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
