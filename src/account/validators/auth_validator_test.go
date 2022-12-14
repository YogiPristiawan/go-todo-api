package validators

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/shared/validators"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthValidatorTestSuite struct {
	suite.Suite
	authValidator AuthValidator
}

func (s *AuthValidatorTestSuite) SetupSuite() {
	s.authValidator = NewAuthValidator(validators.NewValidator())
}

func TestAuthValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(AuthValidatorTestSuite))
}

func (s *AuthValidatorTestSuite) TestNewAuthValidator() {
	s.Run("it should properly instantiate authValidator", func() {
		// arrange
		validator := validators.NewValidator()

		// action
		authValidator := NewAuthValidator(validator)

		// asert
		s.Assert().IsType((AuthValidator)(authValidator), authValidator)
		s.Assert().Implements((*AuthValidator)(nil), authValidator)
	})
}

func (s *AuthValidatorTestSuite) TestValidateLogin() {
	type test struct {
		title     string
		param     dto.LoginRequest
		expectErr *validators.ValidatorError
	}

	s.Run("It should validate username", func() {
		// NEGATIVE
		// arrange
		negative := []test{
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
		for _, test := range negative {
			err := s.authValidator.ValidateLogin(test.param)

			message := "%s"
			message += "\nUsername: %s"
			message += "\nPassword: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message,
				test.title, test.param.Username, test.param.Password,
			)
		}
	})

	s.Run("It should vaildate password", func() {
		// NEGATIVE
		// arrange
		negative := []test{
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
		for _, test := range negative {
			err := s.authValidator.ValidateLogin(test.param)

			message := "%s"
			message += "\nUsername: %s"
			message += "\nPassword: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message,
				test.title, test.param.Username, test.param.Password,
			)
		}
	})

	s.Run("It should return error nil if the given request is valid", func() {
		// arrange
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
			err := s.authValidator.ValidateLogin(test.param)

			message := "%s"
			message += "\nUsername: %s"
			message += "\nPassword: %s"

			s.Assert().Nilf(
				err,
				message,
				test.title, test.param.Username, test.param.Password,
			)
		}
	})
}

func (s *AuthValidatorTestSuite) TestValdiateRegister() {
	// arrange
	type test struct {
		title     string
		param     dto.RegisterRequest
		expectErr *validators.ValidatorError
	}

	s.Run("It should validate username", func() {
		// arrange
		tests := []test{
			{
				title: "Should return an error if the given username is not provided",
				param: dto.RegisterRequest{
					Password:  "12345678",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should return an error if the given username is invalid",
				param: dto.RegisterRequest{
					Username:  "foo123@",
					Password:  "12345678",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
		}

		// action & assert
		for _, test := range tests {
			err := s.authValidator.ValidateRegister(test.param)

			message := "%s"
			message += "\nUsername: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message,
				test.title, test.param.Username,
			)
		}
	})

	s.Run("It should validate password", func() {
		// arrange
		tests := []test{
			{
				title: "Should return an error if password is not provided",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should return an error if password less than 6 characters",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Password:  "123",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
		}

		// action & asert
		for _, test := range tests {
			err := s.authValidator.ValidateRegister(test.param)

			message := "%s"
			message += "\nPassword: %s"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message,
				test.title, test.param.Password,
			)
		}
	})

	s.Run("It should vaildate gender", func() {
		// NEGATIVE TEST
		// arrange
		negatives := []test{
			{
				title: "Should return an error if the given gender is invalid",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Password:  "123456789",
					Gender:    entities.String{Valid: true, String: "U"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should return an error if the given gender is invalid",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Password:  "123456789",
					Gender:    entities.String{Valid: true, String: "p"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should return an error if the given gender is invalid",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Password:  "123456789",
					Gender:    entities.String{Valid: true, String: "l"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
		}

		// action & assert
		for _, test := range negatives {
			err := s.authValidator.ValidateRegister(test.param)

			message := "%s"
			message += "\nGender: %#v"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message,
				test.title, test.param.Gender,
			)
		}

		// POSITIVE TEST
		// arrange
		positives := []test{
			{
				title: "Should not return an error if the given gender is empty",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Password:  "123456789",
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should not return an error if the given gender is valid",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Gender:    entities.String{Valid: true, String: "L"},
					Password:  "123456789",
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should not return an error if the given gender is valid",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Gender:    entities.String{Valid: true, String: "P"},
					Password:  "123456789",
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
			{
				title: "Should not return an error if the given gender is valid",
				param: dto.RegisterRequest{
					Username:  "yuu",
					Gender:    entities.String{Valid: false, String: ""},
					Password:  "123456789",
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
		}

		// action & assert
		for _, test := range positives {
			err := s.authValidator.ValidateRegister(test.param)

			message := "%s"
			message += "\nGender: %#v"

			s.Assert().Nilf(
				err,
				test.title, test.param.Gender,
			)
		}
	})

	s.Run("It should validate birth date", func() {
		// NEGATIVE
		// arrange
		negatives := []test{
			{
				title: "Should return error if the given birth date format is invalid",
				param: dto.RegisterRequest{
					Username:  "foo",
					Password:  "123456778",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "2022"},
				},
			},
			{
				title: "Should return error if the given birth date format is invalid",
				param: dto.RegisterRequest{
					Username:  "foo",
					Password:  "123456778",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "January"},
				},
			},
			{
				title: "Should return error if the given birth date format is invalid",
				param: dto.RegisterRequest{
					Username:  "foo",
					Password:  "123456778",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "06"},
				},
			},
		}

		// action & assert
		for _, test := range negatives {
			err := s.authValidator.ValidateRegister(test.param)

			message := "%s"
			message += "\nBirthDate: %#v"
			message += "\nActualErr: %T"

			s.Assert().ErrorAsf(
				err,
				&test.expectErr,
				message, test.title, test.param.BirthDate, err,
			)
		}

		// POSITIVE
		// arrange
		positives := []test{
			{
				title: "Should not return an error if the given birth date is valid",
				param: dto.RegisterRequest{
					Username:  "foo",
					BirthDate: entities.Date{Valid: true, String: "2022-01-01"},
					Password:  "12345678",
					Gender:    entities.String{Valid: true, String: "L"},
				},
			},
		}

		// action & assert
		for _, test := range positives {
			err := s.authValidator.ValidateRegister(test.param)

			message := "%s"
			message += "\nBirthDate: %#v"
			message += "\nActualErr: %T"

			s.Assert().Nilf(
				err,
				message, test.title, test.param.BirthDate, err,
			)
		}
	})
}
