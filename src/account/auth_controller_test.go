package account

import (
	"go_todo_api/src/account/services/servicesfakes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite

	ctx            *gin.Context
	authService    *servicesfakes.FakeAuthService
	authController AuthController
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}

func (s *AuthControllerTestSuite) SetupTest() {
	// set gin context
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	s.ctx.Set("auth_user_id", int64(1))
	s.ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	s.authService = &servicesfakes.FakeAuthService{}
	s.authController = NewAuthController(s.authService)
}

func (s *AuthControllerTestSuite) TestNewAuthController() {
	s.Run("It should properly instantiate authController", func() {
		// arrange
		var expectedType *authController

		// action
		authController := NewAuthController(s.authService)

		// assert
		s.Assert().Implementsf(
			(*AuthController)(nil),
			authController,
			"authController should implements AuthController interface",
		)
		s.Assert().IsTypef(
			expectedType,
			authController,
			"authController should have the correct type\nGot: %T\nExpect: %T",
			authController, expectedType,
		)
	})
}

func (s *AuthControllerTestSuite) TestLogin() {
	s.Run("It should call the correct service", func() {
		// arrange
		s.SetupTest()

		// action
		s.authController.Login(s.ctx)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.authService.LoginCallCount(),
			"authService.Login has to be called once\nCount: %d\nExpect: %d",
			s.authService.LoginCallCount(), 1,
		)
	})
}

func (s *AuthControllerTestSuite) TestRegister() {
	s.Run("It should call the correct service", func() {
		// arrange
		s.SetupTest()

		// action
		s.authController.Register(s.ctx)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.authService.RegisterCallCount(),
			"authService.Register has to be called once\nCount: %d\nExpect: %d",
			s.authService.RegisterCallCount(), 1,
		)
	})
}
