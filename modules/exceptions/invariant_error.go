package exceptions

import (
	"fmt"
)

type InvariantError struct {
	ClientError
}

func (i *InvariantError) Error() string {
	return fmt.Sprintf("status %d: err %v", i.StatusCode, i.Message)
}

func NewInvariantError(message string) *InvariantError {
	return &InvariantError{
		ClientError{
			Message:    message,
			StatusCode: 400,
		},
	}
}
