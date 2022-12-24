package validators

import (
	"go_todo_api/src/todo/dto"

	"github.com/go-playground/validator/v10"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// TodoValidator is an abstract that contains
// methods to validate todo data
//
//counterfeiter:generate . TodoValidator
type TodoValidator interface {
	ValidateStore(in dto.StoreTodoRequest) error
	ValidateDetail(in dto.DetailTodoRequest) error
	ValidateUpdate(in dto.UpdateTodoRequest) error
}

// todoValidator is a struct that has methods
// to validate todo data
type todoValidator struct {
	validator *validator.Validate
}

// NewTodoValidator creates instance of todoValidator
func NewTodoValidator(validator *validator.Validate) TodoValidator {
	return &todoValidator{
		validator: validator,
	}
}

// ValidateStore handle action to validate
// todo store action data
func (t *todoValidator) ValidateStore(in dto.StoreTodoRequest) error {
	return customErrorMsg(t.validator.Struct(in))
}

// ValidateDetail handle action to validate todo get detail
// action data
func (t *todoValidator) ValidateDetail(in dto.DetailTodoRequest) error {
	return customErrorMsg(t.validator.Struct(in))
}

func (t *todoValidator) ValidateUpdate(in dto.UpdateTodoRequest) error {
	return customErrorMsg(t.validator.Struct(in))
}
