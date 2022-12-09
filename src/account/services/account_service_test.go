package services

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/models"
	"go_todo_api/src/account/repositories/repositoriesfakes"
	"go_todo_api/src/shared/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountService(t *testing.T) {
	t.Run("It should properly instantiate accountService", func(t *testing.T) {
		// arrange
		accountRepository := repositoriesfakes.FakeAccountRepository{}
		var expectedType accountService

		// action
		accountService := NewAccountService(&accountRepository)

		// assert
		assert.Implements(t, (*AccountService)(nil), accountService)
		assert.IsType(t, &expectedType, accountService)
	})
}

func TestGetProfile(t *testing.T) {
	type test struct {
		title  string
		param  dto.ProfileRequest
		expect entities.BaseResponse[dto.ProfileResponse]
	}

	t.Run("It should return response 404 when profile is not found", func(t *testing.T) {
		// arrange
		accountRepository := repositoriesfakes.FakeAccountRepository{}
		accountService := NewAccountService(&accountRepository)

		// mock
		wrapDBErr = func(err error) (code int) {
			val := err.(*mockError)
			return val.code
		}
		accountRepository.GetProfileByUserIdReturns(models.Profile{}, &mockError{
			code: 404,
		})

		negative := test{
			title: "Should provided the correct response",
			param: dto.ProfileRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 1,
				},
			},
			expect: entities.BaseResponse[dto.ProfileResponse]{
				Data: nil,
			},
		}

		// action
		res := accountService.GetProfile(negative.param)

		// assert response message
		message := "%s"
		message += "\nMessage: %s"
		message += "\nExpect: %s"
		assert.Equalf(
			t,
			"profile not found",
			res.GetMessage(),
			message,
			negative.title, res.Message, negative.expect.Message,
		)

		// assert response data
		message = "%s"
		message += "\nData: %#v"
		message += "\nExpect: %#v"
		assert.Exactlyf(
			t,
			negative.expect.Data,
			res.Data,
			message,
			negative.title, res.Data, negative.expect.Data,
		)

		// assert response code
		message = "%s"
		message += "\nCode: %d"
		message += "\nExpect: %d"
		assert.Equalf(
			t,
			404,
			res.GetCode(),
			message,
			negative.title,
			res.GetCode(),
			404,
		)
	})

	t.Run("It should return response 500 when server error is ouccured", func(t *testing.T) {
		// arrange
		accountRepository := repositoriesfakes.FakeAccountRepository{}
		accountService := NewAccountService(&accountRepository)

		// mock
		wrapDBErr = func(err error) (code int) {
			val := err.(*mockError)
			return val.code
		}
		accountRepository.GetProfileByUserIdReturns(models.Profile{}, &mockError{
			code: 500,
		})

		negative := test{
			title: "Should provided the correct response",
			param: dto.ProfileRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 1,
				},
			},
			expect: entities.BaseResponse[dto.ProfileResponse]{
				Data: nil,
			},
		}

		// action
		res := accountService.GetProfile(dto.ProfileRequest{})

		// assert error message
		message := "%s"
		message += "\nMessage: %s"
		message += "\nExpect: %s"

		assert.Exactlyf(
			t,
			"internal server error",
			res.GetMessage(),
			message,
			negative.title, res.Message, negative.expect.Message,
		)
	})

	t.Run("It should return response 200 when error is not occured", func(t *testing.T) {
		// arrange
		accountRepository := repositoriesfakes.FakeAccountRepository{}
		accountService := NewAccountService(&accountRepository)
		positive := test{
			title: "Should provide the correct response",
			param: dto.ProfileRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 1,
				},
			},
			expect: entities.BaseResponse[dto.ProfileResponse]{
				Message: "success get user profile",
				Data: &dto.ProfileResponse{
					Id:        1,
					Username:  "yuu",
					Gender:    entities.String{Valid: true, String: "L"},
					BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
				},
			},
		}

		// mock
		wrapDBErr = func(err error) (code int) {
			return
		}
		accountRepository.GetProfileByUserIdReturns(models.Profile{
			Id:        1,
			Username:  "yuu",
			Gender:    entities.String{Valid: true, String: "L"},
			BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
		}, nil)

		// action
		res := accountService.GetProfile(positive.param)

		// assert response
		message := "%s"
		message += "\nActual: %#v"
		message += "\nExpect: %#v"

		assert.Exactlyf(
			t,
			positive.expect,
			res,
			message,
			positive.title, res, positive.expect,
		)

	})
}
