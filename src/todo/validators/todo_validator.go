package validators

import (
	"go_todo_api/src/shared/validators"
	"go_todo_api/src/todo/dto"
)

// TodoValidator is an abstract that contains
// methods to validate todo data
type TodoValidator interface {
	ValidateStore(in dto.StoreTodoRequest) error
	ValidateDetail(in dto.DetailTodoRequest) error
}

// todoValidator is a struct that has methods
// to validate todo data
type todoValidator struct {
	validator validators.Validate
}

// NewTodoValidator creates instance of todoValidator
func NewTodoValidator(validator validators.Validate) TodoValidator {
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
