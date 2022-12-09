package services

import (
	"fmt"
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/models"
	"go_todo_api/src/account/repositories/repositoriesfakes"
	"go_todo_api/src/account/validators/validatorsfakes"
	"go_todo_api/src/shared/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthService(t *testing.T) {
	t.Run("It should create appropiate instance of authService", func(t *testing.T) {
		// arrange
		authValidator := validatorsfakes.FakeAuthValidator{}
		accountRepository := repositoriesfakes.FakeAccountRepository{}
		var expect authService

		// action
		authService := NewAuthService(&authValidator, &accountRepository)

		// assert
		assert.Implements(t, (*AuthService)(nil), authService)
		assert.IsType(t, &expect, authService)
	})
}

func TestLogin(t *testing.T) {
	type test struct {
		param  dto.LoginRequest
		expect entities.BaseResponse[dto.LoginResponse]
	}

	t.Run("Test http response code 400", func(t *testing.T) {

		t.Run("It should return the correct response if the request validation is failed", func(t *testing.T) {
			// arrange
			authValidator := validatorsfakes.FakeAuthValidator{}
			accountRepository := repositoriesfakes.FakeAccountRepository{}
			authService := NewAuthService(&authValidator, &accountRepository)

			negative := test{
				param: dto.LoginRequest{},
				expect: entities.BaseResponse[dto.LoginResponse]{
					Message: "some validation is failed",
					Data:    nil,
				},
			}

			// mock
			authValidator.ValidateLoginReturns(fmt.Errorf("some validation is failed"))

			// action
			res := authService.Login(negative.param)

			// assert response message
			assert.Equalf(
				t,
				negative.expect.Message,
				res.GetMessage(),
				"should return correct response message\nMessage: %s\nExpect: %s",
				res.Message, negative.expect.Message,
			)

			// assert response data
			assert.Exactlyf(
				t,
				negative.expect.Data,
				res.Data,
				"should return correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			assert.Equalf(
				t,
				400,
				res.GetCode(),
				"should return correct resposne code\nCode: %d\nExpect: %d",
				res.GetCode(), 400,
			)

			// assert method call count
			assert.Equalf(
				t,
				1,
				authValidator.ValidateLoginCallCount(),
				"validation method should called once.\nCount: %d\nExpect: %d",
				authValidator.ValidateLoginCallCount(), 1,
			)
			assert.Equalf(
				t,
				0,
				accountRepository.GetByUsernameCallCount(),
				"GetByUsername should not to be called when validation is failed\nCount: %d\nExpect: %d",
				accountRepository.CreateCallCount(), 0,
			)
		})
	})

	t.Run("Test http response code 401", func(t *testing.T) {

		t.Run("It should return the correct response if the password mismatch", func(t *testing.T) {
			// arrange
			authValidator := validatorsfakes.FakeAuthValidator{}
			accountRepository := repositoriesfakes.FakeAccountRepository{}
			accountService := NewAuthService(&authValidator, &accountRepository)
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
			authValidator.ValidateLoginReturns(nil)
			accountRepository.GetByUsernameReturns(models.Account{
				Id:        1,
				Username:  "yuu",
				Password:  password,
				Gender:    entities.String{Valid: true, String: "L"},
				BirthDate: entities.Date{Valid: true, String: "2022-01-01"},
			}, nil)

			// action
			res := accountService.Login(negative.param)

			// assert response message
			assert.Equalf(
				t,
				negative.expect.Message,
				res.GetMessage(),
				"should return correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			assert.Equalf(
				t,
				negative.expect.Data,
				res.Data,
				"shold return correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			assert.Equalf(
				t,
				401,
				res.GetCode(),
				"should return correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 401,
			)

			// assert method calls count
			assert.Equalf(
				t,
				1,
				authValidator.ValidateLoginCallCount(),
				"validation method should called once\nCount: %d\nExpect: %d",
				authValidator.ValidateLoginCallCount(), 1,
			)
			assert.Equalf(
				t,
				1,
				accountRepository.GetByUsernameCallCount(),
				"GetByUsername should called once\nCount: %d\nExpect: %d",
				accountRepository.GetByUsernameCallCount(), 1,
			)
		})

	})

	t.Run("Test http response code 404", func(t *testing.T) {

		t.Run("It should return the correct response if the username not found", func(t *testing.T) {
			// arrange
			authValidator := validatorsfakes.FakeAuthValidator{}
			accountRepository := repositoriesfakes.FakeAccountRepository{}
			authService := NewAuthService(&authValidator, &accountRepository)

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

			wrapDBErr = func(err error) (code int) {
				val := err.(*mockError)
				return val.code
			}

			// mock
			authValidator.ValidateLoginReturns(nil)
			accountRepository.GetByUsernameReturns(models.Account{}, &mockError{
				code: 404,
			})

			// action
			res := authService.Login(negative.param)

			// assert response message
			assert.Equalf(
				t,
				negative.expect.Message,
				res.GetMessage(),
				"it should have the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			assert.Equalf(
				t,
				negative.expect.Data,
				res.Data,
				"it should have the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			assert.Equalf(
				t,
				404,
				res.GetCode(),
				"it should have the correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 404,
			)

			// assert method call count
			assert.Equalf(
				t,
				1,
				authValidator.ValidateLoginCallCount(),
				"ValidateLoginCallCount should to be called once\nCount: %d\nExpect: %d",
				authValidator.ValidateLoginCallCount(), 1,
			)
			assert.Equalf(
				t,
				1,
				accountRepository.GetByUsernameCallCount(),
				"GetByUsername should to be called once\nCount: %d\nExpect: %d",
				accountRepository.CreateCallCount(), 1,
			)
		})
	})

}
