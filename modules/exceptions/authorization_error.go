package exceptions

import (
	"fmt"
)

type AuthorizationError struct {
	ClientError
}

func (a *AuthorizationError) Error() string {
	return fmt.Sprintf("status %d: err %v", a.StatusCode, a.Message)
}

func NewAuthorizationError(message string) *AuthorizationError {
	return &AuthorizationError{
		ClientError{
			Message:    message,
			StatusCode: 403,
		},
	}
}
