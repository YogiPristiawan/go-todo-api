package services

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/models"
	"go_todo_api/src/account/repositories/repositoriesfakes"
	"go_todo_api/src/shared/entities"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountServiceTestSuite struct {
	suite.Suite
	accountRepository *repositoriesfakes.FakeAccountRepository

	accountService AccountService
}

func TestAccountServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

func (s *AccountServiceTestSuite) SetupSuite() {
	wrapDBErr = func(err error) (code int) {
		if val, ok := err.(*mockError); ok {
			return val.code
		}
		return
	}
}

func (s *AccountServiceTestSuite) SetupTest() {
	s.accountRepository = &repositoriesfakes.FakeAccountRepository{}

	s.accountService = NewAccountService(s.accountRepository)
}

func (s *AccountServiceTestSuite) TestNewAccountService() {
	s.Run("It should properly instantiate accountService", func() {
		s.SetupTest()
		// arrange
		var expectedType accountService

		// action
		accountService := NewAccountService(s.accountRepository)

		// assert
		s.Assert().Implements((*AccountService)(nil), accountService)
		s.Assert().IsType(&expectedType, accountService)
	})
}

func (s *AccountServiceTestSuite) TestGetProfile() {
	type test struct {
		param  dto.ProfileRequest
		expect entities.BaseResponse[dto.ProfileResponse]
	}

	s.Run("Test response code 404", func() {
		s.Run("profile is not found", func() {
			s.SetupTest()
			// mock
			s.accountRepository.GetProfileByUserIdReturns(models.Profile{}, &mockError{
				code: 404,
			})

			negative := test{
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
			res := s.accountService.GetProfile(negative.param)

			// assert response message
			s.Assert().Equalf(
				"profile not found",
				res.GetMessage(),
				"it should return the correct response message\nMessage: %s\nExpect: %s",
				res.Message, negative.expect.Message,
			)

			// assert response data
			s.Assert().Exactlyf(
				negative.expect.Data,
				res.Data,
				"It should return the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				404,
				res.GetCode(),
				"It should return the correct response code\nCode: %s\nExpect: %d",
				res.GetCode(),
				404,
			)
		})
	})

	s.Run("Test response code 500", func() {
		s.SetupTest()
		// mock
		s.accountRepository.GetProfileByUserIdReturns(models.Profile{}, &mockError{
			code: 500,
		})

		negative := test{
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
		res := s.accountService.GetProfile(negative.param)

		// assert error code
		s.Assert().Exactlyf(
			500,
			res.GetCode(),
			"It should return correct response code\nCode: %d\nExpect: %d",
			res.GetCode(), 500,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.accountRepository.GetProfileByUserIdCallCount(),
			"GetProfileByUserId has to be called once\nCount: %d\nExpect: %d",
			s.accountRepository.GetProfileByUserIdCallCount(), 1,
		)
	})

	s.Run("Test response code 200", func() {
		s.SetupTest()
		// arrange
		positive := test{
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
		s.accountRepository.GetProfileByUserIdReturns(models.Profile{
			Id:        1,
			Username:  "yuu",
			Gender:    entities.String{Valid: true, String: "L"},
			BirthDate: entities.Date{Valid: true, String: "2000-01-01"},
		}, nil)

		// action
		res := s.accountService.GetProfile(positive.param)

		// assert response
		s.Assert().Exactlyf(
			positive.expect,
			res,
			"it sould return correct response\nResponse: %#v\nExpect: %#v",
			res, positive.expect,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.accountRepository.GetProfileByUserIdCallCount(),
			"GetProfileByUserId has to be called once\nCount: %d\nExpect: %d",
			s.accountRepository.GetProfileByUserIdCallCount(), 1,
		)

	})
}
