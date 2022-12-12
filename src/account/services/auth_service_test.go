package services

import (
	"fmt"
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/models"
	"go_todo_api/src/account/repositories/repositoriesfakes"
	"go_todo_api/src/account/validators/validatorsfakes"
	"go_todo_api/src/shared/entities"

	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	authValidator     *validatorsfakes.FakeAuthValidator
	accountRepository *repositoriesfakes.FakeAccountRepository

	authService AuthService
}

func (s *AuthServiceTestSuite) SetupSuite() {
	wrapDBErr = func(err error) (code int) {
		if val, ok := err.(*mockError); ok {
			return val.code
		}
		return
	}
}

func (s *AuthServiceTestSuite) SetupTest() {
	s.authValidator = &validatorsfakes.FakeAuthValidator{}
	s.accountRepository = &repositoriesfakes.FakeAccountRepository{}

	s.authService = NewAuthService(s.authValidator, s.accountRepository)
}

func (s *AuthServiceTestSuite) TestNewAuthService() {
	s.Run("It should create appropiate instance of authService", func() {
		// arrange
		authValidator := validatorsfakes.FakeAuthValidator{}
		accountRepository := repositoriesfakes.FakeAccountRepository{}
		var expect authService

		// action
		authService := NewAuthService(&authValidator, &accountRepository)

		// assert
		s.Assert().Implements((*AuthService)(nil), authService)
		s.Assert().IsType(&expect, authService)
	})
}

func (s *AuthServiceTestSuite) TestLogin() {
	type test struct {
		param  dto.LoginRequest
		expect entities.BaseResponse[dto.LoginResponse]
	}

	s.Run("Test http response code 400", func() {

		s.Run("validation failed", func() {
			s.SetupTest()
			// arrange
			negative := test{
				param: dto.LoginRequest{},
				expect: entities.BaseResponse[dto.LoginResponse]{
					Message: "some validation is failed",
					Data:    nil,
				},
			}

			// mock
			s.authValidator.ValidateLoginReturns(fmt.Errorf("some validation is failed"))

			// action
			res := s.authService.Login(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"should return correct response message\nMessage: %s\nExpect: %s",
				res.Message, negative.expect.Message,
			)

			// assert response data
			s.Assert().Exactlyf(
				negative.expect.Data,
				res.Data,
				"should return correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				400,
				res.GetCode(),
				"should return correct resposne code\nCode: %d\nExpect: %d",
				res.GetCode(), 400,
			)

			// assert method call count
			s.Assert().Equalf(
				1,
				s.authValidator.ValidateLoginCallCount(),
				"validation method should called once.\nCount: %d\nExpect: %d",
				s.authValidator.ValidateLoginCallCount(), 1,
			)
			s.Assert().Equalf(
				0,
				s.accountRepository.GetByUsernameCallCount(),
				"GetByUsername should not to be called when validation is failed\nCount: %d\nExpect: %d",
				s.accountRepository.CreateCallCount(), 0,
			)
		})
	})

	s.Run("Test http response code 401", func() {

		s.Run("Password mismatch", func() {
			s.SetupTest()
			// arrange
			var password = "$2a$12$jhTVu0HPAzwDEbEkC/GgmeMoRCNymiRHNtgVn2RdjSQHcsWY0BNSu" // 12345678

			negative := test{
				param: dto.LoginRequest{
					Username: "yuu",
					Password: "qwerrtyuiop",
				},
				expect: entities.BaseResponse[dto.LoginResponse]{
					Message: "password tidak sesuai",
					Data:    nil,
				},
			}

			// mock
			s.authValidator.ValidateLoginReturns(nil)
			s.accountRepository.GetByUsernameReturns(models.Account{
				Id:        1,
				Username:  "yuu",
				Password:  password,
				Gender:    entities.String{Valid: true, String: "L"},
				BirthDate: entities.Date{Valid: true, String: "2022-01-01"},
			}, nil)

			// action
			res := s.authService.Login(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"should return correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"shold return correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				401,
				res.GetCode(),
				"should return correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 401,
			)

			// assert method calls count
			s.Assert().Equalf(
				1,
				s.authValidator.ValidateLoginCallCount(),
				"validation method should called once\nCount: %d\nExpect: %d",
				s.authValidator.ValidateLoginCallCount(), 1,
			)
			s.Assert().Equalf(
				1,
				s.accountRepository.GetByUsernameCallCount(),
				"GetByUsername should called once\nCount: %d\nExpect: %d",
				s.accountRepository.GetByUsernameCallCount(), 1,
			)
		})

	})

	s.Run("Test http response code 404", func() {

		s.Run("It should return the correct response if the username not found", func() {
			s.SetupTest()
			// arrange
			negative := test{
				param: dto.LoginRequest{
					Username: "yuu",
					Password: "12345678",
				},
				expect: entities.BaseResponse[dto.LoginResponse]{
					Message: "username tidak ditemukan",
					Data:    nil,
				},
			}

			// mock
			s.authValidator.ValidateLoginReturns(nil)
			s.accountRepository.GetByUsernameReturns(models.Account{}, &mockError{
				code: 404,
			})

			// action
			res := s.authService.Login(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"it should have the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"it should have the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				404,
				res.GetCode(),
				"it should have the correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 404,
			)

			// assert method calls count
			s.Assert().Equalf(
				1,
				s.authValidator.ValidateLoginCallCount(),
				"ValidateLoginCallCount should to be called once\nCount: %d\nExpect: %d",
				s.authValidator.ValidateLoginCallCount(), 1,
			)
			s.Assert().Equalf(
				1,
				s.accountRepository.GetByUsernameCallCount(),
				"GetByUsername should to be called once\nCount: %d\nExpect: %d",
				s.accountRepository.CreateCallCount(), 1,
			)
		})
	})

}
