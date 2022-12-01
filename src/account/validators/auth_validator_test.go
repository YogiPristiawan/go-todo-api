package validators

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/shared/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthValidator(t *testing.T) {
	t.Run("it should properly instantiate authValidator", func(t *testing.T) {
		// arrange
		validator := validators.NewValidator()

		// action
		authValidator := NewAuthValidator(validator)

		// asert
		assert.IsType(t, (AuthValidator)(authValidator), authValidator)
		assert.Implements(t, (*AuthValidator)(nil), authValidator)
	})
}

func TestValidateLogin(t *testing.T) {
	type test struct {
		title     string
		param     dto.LoginRequest
		expectErr *validators.ValidatorError
	}

	// arange
	validator := validators.NewValidator()
	authValidator := NewAuthValidator(validator)

	t.Run("It Should validate username", func(t *testing.T) {
		tests := []test{
			{
				title: "Should return an error when username is not provided",
				param: dto.LoginRequest{
					Password: "12345678",
				},
			},
			{
				title: "Should return error when the given username contains prohibitted character",
				param: dto.LoginRequest{
					Username: "testing=@#41234",
					Password: "12345678",
				},
			},
		}

		// assert
		for _, test := range tests {
			err := authValidator.ValidateLogin(test.param)

			message := "%s"
			message += "\nUsername: %s"
			message += "\nPassword: %s"

			assert.ErrorAsf(t,
				err,
				&test.expectErr,
				message,
				test.title, test.param.Username, test.param.Password,
			)
		}
	})

	t.Run("It should vaildate password", func(t *testing.T) {
		// arrange
		tests := []test{
			{
				title: "Should return an error if password is not provided",
				param: dto.LoginRequest{
					Username: "yuu",
				},
			},
			{
				title: "Should return an error if password length less than 6 characters",
				param: dto.LoginRequest{
					Username: "yuu",
					Password: "123",
				},
			},
		}

		// action & assert
		for _, test := range tests {
			err := authValidator.ValidateLogin(test.param)

			message := "%s"
			message += "\nUsername: %s"
			message += "\nPassword: %s"

			assert.ErrorAsf(
				t,
				err,
				&test.expectErr,
				message,
				test.title, test.param.Username, test.param.Password,
			)
		}
	})

	t.Run("It should return error nil if the given request is valid", func(t *testing.T) {
		// arrang
		tests := []test{
			{
				title: "Should not return an error",
				param: dto.LoginRequest{
					Username: "foo",
					Password: "1234567890",
				},
			},
			{
				title: "Should not return an error",
				param: dto.LoginRequest{
					Username: "foo",
					Password: "123ABC@#$=",
				},
			},
		}

		// action & assert
		for _, test := range tests {
			err := authValidator.ValidateLogin(test.param)

			message := "%s"
			message += "\nUsername: %s"
			message += "\nPassword: %s"

			assert.Nilf(
				t,
				err,
				message,
				test.title, test.param.Username, test.param.Password,
			)
		}
	})
}
