package exceptions

import (
	"fmt"
)

type NotFoundError struct {
	ClientError
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("status %d: err %v", n.StatusCode, n.Message)
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		ClientError{
			Message:    message,
			StatusCode: 404,
		},
	}
}
