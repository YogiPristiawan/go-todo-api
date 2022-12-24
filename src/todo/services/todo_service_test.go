package services

import (
	"fmt"
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/todo/dto"
	"go_todo_api/src/todo/models"
	"go_todo_api/src/todo/repositories/repositoriesfakes"
	"go_todo_api/src/todo/validators/validatorsfakes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TodoServiceTestSuite struct {
	suite.Suite
	todoService    TodoService
	todoValidator  *validatorsfakes.FakeTodoValidator
	todoRepository *repositoriesfakes.FakeTodoRepository
}

func TestTodoServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TodoServiceTestSuite))
}

func (s *TodoServiceTestSuite) SetupSuite() {
	wrapDBErr = func(err error) (code int) {
		if val, ok := err.(*mockError); ok {
			return val.code
		}
		return
	}
}

func (s *TodoServiceTestSuite) SetupTest() {
	s.todoValidator = &validatorsfakes.FakeTodoValidator{}
	s.todoRepository = &repositoriesfakes.FakeTodoRepository{}

	s.todoService = NewTodoService(s.todoValidator, s.todoRepository)
}

func (s *TodoServiceTestSuite) TestNewTodoService() {
	// arrange
	var expectType todoService

	// action
	todoService := NewTodoService(s.todoValidator, s.todoRepository)

	// assert
	s.Assert().Implementsf((*TodoService)(nil), todoService, "should implement the correct interface")
	s.Assert().IsType(&expectType, todoService, "should has the correct type")
}

func (s *TodoServiceTestSuite) TestStore() {
	type test struct {
		param  dto.StoreTodoRequest
		expect entities.BaseResponse[dto.StoreTodoResponse]
	}

	s.Run("Test response code 201", func() {
		s.SetupTest()
		// arrange
		positive := test{
			param: dto.StoreTodoRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 2,
				},
				Todo:       "this is a todo",
				Date:       "2022-01-01",
				IsFinished: false,
			},
			expect: entities.BaseResponse[dto.StoreTodoResponse]{
				Message: "todo created",
				Data: &dto.StoreTodoResponse{
					Id:         1,
					UserId:     2,
					Todo:       "this is a todo",
					Date:       "2022-01-01",
					IsFinished: false,
					CreatedAt:  12345678,
					UpdatedAt:  123456778,
				},
			},
		}

		// mock
		s.todoRepository.StoreCalls(func(model *models.Todo) error {
			if model == nil {
				return &mockError{code: 500}
			}
			model.Id = positive.expect.Data.Id
			model.CreatedAt = positive.expect.Data.CreatedAt
			model.UpdatedAt = positive.expect.Data.UpdatedAt
			return nil
		})

		// action
		res := s.todoService.Store(positive.param)

		// assert response message
		s.Assert().Equalf(
			positive.expect.Message,
			res.GetMessage(),
			"should return the correct response message\nMessage: %s\nExpect: %s",
			res.GetMessage(), positive.expect.Message,
		)

		// assert response data
		s.Assert().Exactlyf(
			positive.expect.Data,
			res.Data,
			"should return the correct response data\nData: %#v\nExpect: %#v",
			res.Data, positive.expect.Data,
		)

		// assert response code
		s.Assert().Equalf(
			res.GetCode(),
			201,
			"should return the correct response code\nCode: %s\nExpect: %s",
			res.GetCode(), 201,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.todoValidator.ValidateStoreCallCount(),
			"ValidateStore has to be called once\nCount: %d\nExpect: %d",
			s.todoValidator.ValidateDetailCallCount(), 1,
		)
		s.Assert().Equalf(
			1,
			s.todoRepository.StoreCallCount(),
			"Store has to be called once\nCount: %d\nExpect: %d",
			s.todoRepository.StoreCallCount(), 1,
		)
	})

	s.Run("Test response code 400", func() {

		s.Run("it should return correct response when request validation is failed", func() {
			s.SetupTest()
			// arrange
			negative := test{
				param: dto.StoreTodoRequest{},
				expect: entities.BaseResponse[dto.StoreTodoResponse]{
					Message: "some validation is failed",
					Data:    nil,
				},
			}

			// mock
			s.todoValidator.ValidateStoreReturns(fmt.Errorf("some validation is failed"))

			// action
			res := s.todoService.Store(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"should have the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"response data should have the nil value\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				400,
				res.GetCode(),
				"should have the correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 400,
			)

			// assert method calls count
			s.Assert().Equalf(
				s.todoValidator.ValidateStoreCallCount(),
				1,
				"ValidateStore has to be called once\nCount: %d\nExpect: %d",
				s.todoValidator.ValidateDetailCallCount(), 1,
			)
			s.Assert().Equalf(
				s.todoRepository.StoreCallCount(),
				0,
				"StoreCallCount should no be called if request validation is failed\nCount: %d\nExpect: %d",
				s.todoRepository.StoreCallCount(), 0,
			)
		})
	})

	s.Run("Test response 500", func() {
		s.SetupTest()
		// arrange
		negative := test{
			param: dto.StoreTodoRequest{},
		}

		// mock
		s.todoValidator.ValidateStoreReturns(nil)
		s.todoRepository.StoreReturns(&mockError{
			code: 500,
		})

		// action
		res := s.todoService.Store(negative.param)

		// assert response code
		s.Assert().Equalf(
			res.GetCode(),
			500,
			"should return response code 500 if it has server error\nCode: %d\nExpect: %d",
			res.GetCode(), 500,
		)
	})
}

func (s *TodoServiceTestSuite) TestFind() {
	type test struct {
		param  dto.FindTodoRequest
		expect entities.BaseResponseArray[dto.FindTodoResponse]
	}

	s.Run("Test response code 200", func() {
		s.SetupTest()
		// arrange
		positive := test{
			param: dto.FindTodoRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 2,
				},
			},
			expect: entities.BaseResponseArray[dto.FindTodoResponse]{
				Message: "list of todos",
				Data: []dto.FindTodoResponse{
					{
						Id:         1,
						UserId:     2,
						Todo:       "this is a todo 1",
						Date:       "2022-01-01",
						IsFinished: false,
						CreatedAt:  12345678,
						UpdatedAt:  12345678,
					},
					{
						Id:         2,
						UserId:     2,
						Todo:       "this is a todo 2",
						Date:       "2022-02-02",
						IsFinished: true,
						CreatedAt:  123456789,
						UpdatedAt:  123456789,
					},
				},
			},
		}

		// mock
		s.todoRepository.FindReturns([]models.Todo{
			{
				Id:         1,
				UserId:     2,
				Todo:       "this is a todo 1",
				Date:       "2022-01-01",
				IsFinished: false,
				CreatedAt:  12345678,
				UpdatedAt:  12345678,
			},
			{
				Id:         2,
				UserId:     2,
				Todo:       "this is a todo 2",
				Date:       "2022-02-02",
				IsFinished: true,
				CreatedAt:  123456789,
				UpdatedAt:  123456789,
			},
		}, nil)

		// action
		res := s.todoService.Find(positive.param)

		// assert response message
		s.Assert().Equalf(
			positive.expect.Message,
			res.GetMessage(),
			"should return the correct response message\nMessage: %s\nExpect: %s",
			res.GetMessage(), positive.expect.Message,
		)

		// assert response data (shold has the correct orders to)
		s.Assert().Exactlyf(
			positive.expect.Data,
			res.Data,
			"it should return the exactly corrent response data\nData: %#v\nExpect: %#v",
			res.Data, positive.expect.Data,
		)

		// assert response code
		s.Assert().Equalf(
			200,
			res.GetCode(),
			"it should return the correct response code\nCode: %s\nExpect: %d",
			res.GetCode(), 200,
		)
	})

	s.Run("Test response code 404", func() {

		s.Run("it should return correct response when todo is not found", func() {
			s.SetupTest()
			// arrange
			negative := test{
				param: dto.FindTodoRequest{
					RequestMetaData: entities.RequestMetaData{
						AuthUserId: 2,
					},
				},
				expect: entities.BaseResponseArray[dto.FindTodoResponse]{
					Message: "todo not found",
					Data:    nil,
				},
			}

			// mock
			s.todoRepository.FindReturns([]models.Todo{}, &mockError{
				code: 404,
			})

			// action
			res := s.todoService.Find(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"it should return the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"it should return the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				404,
				res.GetCode(),
				"it should return the correct response code\nCode: %s\nExpect: %d",
				res.GetCode(), 404,
			)

			// assert method calls count
			s.Assert().Equalf(
				1,
				s.todoRepository.FindCallCount(),
				"Find methods has to be called once\nCount: %d\nExpect: %d",
				s.todoRepository.FindCallCount(), 1,
			)
		})
	})

	s.Run("Test response code 500", func() {
		s.SetupTest()
		// arrange
		negative := test{
			param: dto.FindTodoRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 2,
				},
			},
			expect: entities.BaseResponseArray[dto.FindTodoResponse]{
				Data: nil,
			},
		}

		// mock
		s.todoRepository.FindReturns([]models.Todo{}, &mockError{
			code: 500,
		})

		// action
		res := s.todoService.Find(negative.param)

		// assert response data
		s.Assert().Equalf(
			negative.expect.Data,
			res.Data,
			"it should return the correct response data\nData: %#v\nExpecct: %#v",
			res.Data, negative.expect.Data,
		)

		// assert response code
		s.Assert().Equalf(
			500,
			res.GetCode(),
			"it should return the correct response code\nCode: %d\nExpect: %d",
			res.GetCode(), 500,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.todoRepository.FindCallCount(),
			"Find method has to be called once\nCount: %d\nExpect: %d",
			s.todoRepository.FindCallCount(), 1,
		)
	})
}

func (s *TodoServiceTestSuite) TestDetail() {
	type test struct {
		param  dto.DetailTodoRequest
		expect entities.BaseResponse[dto.DetailTodoResponse]
	}

	s.Run("Test response code 200", func() {
		s.SetupTest()
		// arrange
		positive := test{
			param: dto.DetailTodoRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 2,
				},
				Id: 1,
			},
			expect: entities.BaseResponse[dto.DetailTodoResponse]{
				Message: "detail of todo",
				Data: &dto.DetailTodoResponse{
					Id:         1,
					UserId:     2,
					Todo:       "this is a todo",
					Date:       "2022-02-02",
					IsFinished: false,
					CreatedAt:  12345678,
					UpdatedAt:  12345678,
				},
			},
		}

		// mock
		s.todoValidator.ValidateDetailReturns(nil)
		s.todoRepository.DetailReturns(models.Todo{
			Id:         1,
			UserId:     2,
			Todo:       "this is a todo",
			Date:       "2022-02-02",
			IsFinished: false,
			CreatedAt:  12345678,
			UpdatedAt:  12345678,
		}, nil)

		// action
		res := s.todoService.Detail(positive.param)

		// assert response message
		s.Assert().Equalf(
			positive.expect.Message,
			res.GetMessage(),
			"it should return the correct response message\nMessage: %s\nExpect: %s",
			res.GetMessage(), positive.expect.Message,
		)

		// assert response data
		s.Assert().Equalf(
			positive.expect.Data,
			res.Data,
			"it should return the correct response data\nData: %#v\nExpect: %#v",
			res.Data, positive.expect.Data,
		)

		// assert response code
		s.Assert().Equalf(
			200,
			res.GetCode(),
			"it should return the correct response code\nCode: %s\nExpect: %d",
			res.GetCode(), 200,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.todoRepository.DetailCallCount(),
			"Detail method has to be called once\nCount: %d\nExpect: %d",
			s.todoRepository.DetailCallCount(), 1,
		)
		s.Assert().Equalf(
			1,
			s.todoValidator.ValidateDetailCallCount(),
			"ValidateDetail method has to be called once\nCount: %d\nExpect: %d",
			s.todoValidator.ValidateDetailCallCount(), 1,
		)
	})

	s.Run("Test response code 400", func() {

		s.Run("Request validation is failed", func() {
			s.SetupTest()
			// arrange
			negative := test{
				param: dto.DetailTodoRequest{
					RequestMetaData: entities.RequestMetaData{
						AuthUserId: 2,
					},
				},
				expect: entities.BaseResponse[dto.DetailTodoResponse]{
					Message: "some validation is failed",
					Data:    nil,
				},
			}

			// mock
			s.todoValidator.ValidateDetailReturns(fmt.Errorf("some validation is failed"))

			// action
			res := s.todoService.Detail(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"it should return the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"it should return the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				400,
				res.GetCode(),
				"it should return the correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 400,
			)

			// assert method calls count
			s.Assert().Equalf(
				1,
				s.todoValidator.ValidateDetailCallCount(),
				"ValidateDetail has to be called once\nCount: %d\nExpect: %d",
				s.todoValidator.ValidateDetailCallCount(), 1,
			)
			s.Assert().Equalf(
				0,
				s.todoRepository.DetailCallCount(),
				"Detail method should not be called if the request valdation is failed\nCount: %d\nExpect: %d",
				s.todoRepository.DetailCallCount(), 0,
			)
		})
	})

	s.Run("Test response code 404", func() {

		s.Run("Todo is not found", func() {
			s.SetupTest()
			// arrange
			negative := test{
				param: dto.DetailTodoRequest{
					RequestMetaData: entities.RequestMetaData{
						AuthUserId: 2,
					},
					Id: 5,
				},
				expect: entities.BaseResponse[dto.DetailTodoResponse]{
					Message: "todo is not found",
					Data:    nil,
				},
			}

			// mock
			s.todoValidator.ValidateDetailReturns(nil)
			s.todoRepository.DetailReturns(models.Todo{}, &mockError{code: 404})

			// action
			res := s.todoService.Detail(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"it should return the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"it should return the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				404,
				res.GetCode(),
				"it should return the correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 404,
			)

			// assert method calls count
			s.Assert().Equalf(
				1,
				s.todoValidator.ValidateDetailCallCount(),
				"ValidateDetail has to be called once\nCount: %d\nExpect: %d",
				s.todoValidator.ValidateDetailCallCount(), 1,
			)
			s.Assert().Equalf(
				1,
				s.todoRepository.DetailCallCount(),
				"Detail method has to be called once\nCount: %d\nExpect: %d",
				s.todoRepository.DetailCallCount(), 1,
			)
		})
	})

	s.Run("Test response code 500", func() {
		s.SetupTest()
		// arrange
		negative := test{
			param: dto.DetailTodoRequest{},
			expect: entities.BaseResponse[dto.DetailTodoResponse]{
				Data: nil,
			},
		}

		// mock
		s.todoValidator.ValidateDetailReturns(nil)
		s.todoRepository.DetailReturns(models.Todo{}, &mockError{code: 500})

		// action
		res := s.todoService.Detail(negative.param)

		// assert response data
		s.Assert().Equalf(
			negative.expect.Data,
			res.Data,
			"it should return the correct response data\nData: %#v\nExpect: %#v",
			res.Data, negative.expect.Data,
		)

		// assert resonse code
		s.Assert().Equalf(
			500,
			res.GetCode(),
			"it should return the correct response code\nCode: %d\nExpect: %d",
			res.GetCode(), 500,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.todoValidator.ValidateDetailCallCount(),
			"ValidateDetail has to be called once\nCount: %d\nExpect: %d",
			s.todoValidator.ValidateDetailCallCount(), 1,
		)
		s.Assert().Equalf(
			1,
			s.todoRepository.DetailCallCount(),
			"Detail has to be called once\nCount: %d\nExpect: %d",
			s.todoRepository.DetailCallCount(), 1,
		)
	})
}

func (s *TodoServiceTestSuite) TestUpdate() {
	type test struct {
		param  dto.UpdateTodoRequest
		expect entities.BaseResponse[dto.UpdateTodoResponse]
	}

	s.Run("Test response code 200", func() {
		// arrange
		s.SetupTest()

		positive := test{
			param: dto.UpdateTodoRequest{
				RequestMetaData: entities.RequestMetaData{
					AuthUserId: 2,
				},
				Id:         1,
				Todo:       "Todo after update",
				Date:       "2022-02-02",
				IsFinished: true,
			},
			expect: entities.BaseResponse[dto.UpdateTodoResponse]{
				Message: "todo updated",
				Data: &dto.UpdateTodoResponse{
					Id:         1,
					Todo:       "Todo after update",
					Date:       "2022-02-02",
					IsFinished: true,
					CreatedAt:  1670874430,
					UpdatedAt:  1671072430,
				},
			},
		}

		// mock
		s.todoRepository.UpdateCalls(func(model *models.Todo) (todo models.Todo, err error) {
			if model == nil {
				err = &mockError{code: 500}
				return
			}
			todo.Id = model.Id
			todo.UserId = model.UserId
			todo.Todo = model.Todo
			todo.Date = model.Date
			todo.IsFinished = model.IsFinished
			todo.CreatedAt = 1670874430
			todo.UpdatedAt = 1671072430
			return
		})

		// action
		res := s.todoService.Update(positive.param)

		// assert response message
		s.Assert().Equalf(
			positive.expect.Message,
			res.GetMessage(),
			"It should return the correct response message\nMessage: %s\nExpect: %s",
			res.GetMessage(), positive.expect.Message,
		)

		// assert response data
		s.Assert().Exactlyf(
			positive.expect.Data,
			res.Data,
			"It should return the correct response data\nData: %#v\nExpect: %#v",
			res.Data, positive.expect.Data,
		)

		// assert response code
		s.Assert().Equalf(
			200,
			res.GetCode(),
			"It should return the correct response code\nCode: %d\nExpect: %d",
			res.GetCode(), 200,
		)

		// assert method calls count
		s.Assert().Equalf(
			1,
			s.todoRepository.UpdateCallCount(),
			"todoRepository.Update has to be called once\nCount: %d\nExpect: %d",
			s.todoRepository.DetailCallCount(), 1,
		)
	})

	s.Run("Test response code 400", func() {
		s.Run("request validation is failed", func() {
			// arrange
			s.SetupTest()

			negative := test{
				param: dto.UpdateTodoRequest{
					RequestMetaData: entities.RequestMetaData{
						AuthUserId: 2,
					},
				},
				expect: entities.BaseResponse[dto.UpdateTodoResponse]{
					Message: "some request validation is failed",
					Data:    nil,
				},
			}

			// mock
			s.todoValidator.ValidateUpdateReturns(fmt.Errorf("some request validation is failed"))

			// action
			res := s.todoService.Update(negative.param)

			// assert response message
			s.Assert().Equalf(
				negative.expect.Message,
				res.GetMessage(),
				"It should return the correct response message\nMessage: %s\nExpect: %s",
				res.GetMessage(), negative.expect.Message,
			)

			// assert response data
			s.Assert().Equalf(
				negative.expect.Data,
				res.Data,
				"It should return the correct response data\nData: %#v\nExpect: %#v",
				res.Data, negative.expect.Data,
			)

			// assert response code
			s.Assert().Equalf(
				400,
				res.GetCode(),
				"It should return the correct response code\nCode: %d\nExpect: %d",
				res.GetCode(), 400,
			)

			// assert method calls count
			s.Assert().Equalf(
				0,
				s.todoRepository.UpdateCallCount(),
				"todoRepository.Update should not be called when validation is failed\nCount: %d\nExpect: %d",
				s.todoRepository.DetailCallCount(), 0,
			)
		})
	})

	s.Run("Test response code 500", func() {
		// arrange
		s.SetupTest()
		negative := test{
			param: dto.UpdateTodoRequest{},
			expect: entities.BaseResponse[dto.UpdateTodoResponse]{
				Data: nil,
			},
		}

		// mock
		s.todoValidator.ValidateUpdateReturns(nil)
		s.todoRepository.UpdateReturns(models.Todo{}, &mockError{code: 500})

		// action
		res := s.todoService.Update(negative.param)

		// assert response data
		s.Assert().Equalf(
			negative.expect.Data,
			res.Data,
			"It should return the correct response data\nData: %#v\nExpect: %#v",
			res.Data, negative.expect.Data,
		)

		// assert response code
		s.Assert().Equalf(
			500,
			res.GetCode(),
			"It should return the correct response code\nCode: %d\nExpect: %d",
			res.GetCode(), 500,
		)
	})
}
