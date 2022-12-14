package todo

import (
	"go_todo_api/src/todo/services/servicesfakes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TodoControllerTestSuite struct {
	suite.Suite
	ctx            *gin.Context
	todoService    *servicesfakes.FakeTodoService
	todoController TodoController
}

func TestTodoControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TodoControllerTestSuite))
}

func (s *TodoControllerTestSuite) SetupTest() {
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	s.ctx.Set("auth_user_id", int64(1))
	s.ctx.Request = &http.Request{
		Header:     make(http.Header),
		RequestURI: "/todos/:2",
	}

	s.todoService = &servicesfakes.FakeTodoService{}
	s.todoController = NewTodoController(s.todoService)
}

func (s *TodoControllerTestSuite) TestNewTodoController() {
	s.Run("It should properly instantiate todoController", func() {
		// arrange
		var expectedType *todoController

		// action
		todoController := NewTodoController(s.todoService)

		// assert
		s.Assert().Implementsf(
			(*TodoController)(nil),
			todoController,
			"todoController should implements TodoController",
		)
		s.Assert().IsTypef(
			expectedType,
			todoController,
			"todoController should have the correct type\nGot: %T\nExpect: %T",
			todoController, expectedType,
		)
	})
}

func (s *TodoControllerTestSuite) TestStore() {
	s.Run("It should calls the correct service", func() {
		// arrange
		s.SetupTest()

		// action
		s.todoController.Store(s.ctx)

		// assert
		s.Assert().Equalf(
			1,
			s.todoService.StoreCallCount(),
			"todoService.Store has to be called once\nCount: %d\nExpect: %d",
			s.todoService.StoreCallCount(), 1,
		)
	})
}

func (s *TodoControllerTestSuite) TestFind() {
	s.Run("It shuld calls the correct service", func() {
		// arrange
		s.SetupTest()

		// action
		s.todoController.Find(s.ctx)

		// assert
		s.Assert().Equalf(
			1,
			s.todoService.FindCallCount(),
			"todoService.Find has to be called once\nCount: %d\nExpect: %d",
			s.todoService.FindCallCount(), 1,
		)
	})
}
