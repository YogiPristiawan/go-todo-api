package account

import (
	"go_todo_api/src/account/services/servicesfakes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AccountControllerTestSuite struct {
	suite.Suite
	ctx            *gin.Context
	accountService *servicesfakes.FakeAccountService
	authService    *servicesfakes.FakeAuthService

	accountController AccountController
}

func TestAccountControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AccountControllerTestSuite))
}

func (s *AccountControllerTestSuite) SetupTest() {
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	s.ctx.Set("auth_user_id", int64(1))

	s.accountService = &servicesfakes.FakeAccountService{}
	s.authService = &servicesfakes.FakeAuthService{}

	s.accountController = NewAccountController(s.accountService)
}

func (s *AccountControllerTestSuite) TestNewAccountController() {
	s.Run("It should properly instantiate accountController", func() {
		// arrange
		var expectedType *accountController

		// action
		accountController := NewAccountController(s.accountService)

		// assert
		s.Assert().Implementsf(
			(*AccountController)(nil),
			accountController,
			"accountController should implements AccountController interface",
		)
		s.Assert().IsTypef(
			expectedType,
			accountController,
			"returned value should have the correct type\nGot: %T\nExpect: %T",
			accountController, expectedType,
		)
	})
}

func (s *AccountControllerTestSuite) TestGetProfile() {
	s.Run("It should calls correct service", func() {
		// action
		s.accountController.GetProfile(s.ctx)

		// assert
		s.Assert().Equalf(
			1,
			s.accountService.GetProfileCallCount(),
			"accountService.GetProfile has to be called once\nCount: %d\nExpect: %d",
			s.accountService.GetProfileCallCount(), 1,
		)
	})
}
