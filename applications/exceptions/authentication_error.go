package exceptions

import (
	"fmt"
)

type AuthenticationError struct {
	ClientError
}

func (a *AuthenticationError) Error() string {
	return fmt.Sprintf("status %d: err %v", a.StatusCode, a.Message)
}

func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{
		ClientError{
			Message:    message,
			StatusCode: 401,
		},
	}
}
