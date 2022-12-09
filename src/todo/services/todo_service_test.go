package services

import (
	"fmt"
	"go_todo_api/src/shared/entities"
	"go_todo_api/src/todo/dto"
	"go_todo_api/src/todo/models"
	"go_todo_api/src/todo/repositories/repositoriesfakes"
	"go_todo_api/src/todo/validators/validatorsfakes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodoService(t *testing.T) {
	// arrange
	todoValidator := validatorsfakes.FakeTodoValidator{}
	todoRepository := repositoriesfakes.FakeTodoRepository{}
	var expectType todoService

	// action
	todoService := NewTodoService(&todoValidator, &todoRepository)

	// assert
	assert.Implementsf(t, (*TodoService)(nil), todoService, "should implement the correct interface")
	assert.IsType(t, &expectType, todoService, "should has the correct type")
}

func TestStore(t *testing.T) {
	type test struct {
		param  dto.StoreTodoRequest
		expect entities.BaseResponse[dto.StoreTodoResponse]
	}

	t.Run("Test response code 201", func(t *testing.T) {
		// arrange
		todoValidator := validatorsfakes.FakeTodoValidator{}
		todoRepository := repositoriesfakes.FakeTodoRepository{}
		todoService := NewTodoService(&todoValidator, &todoRepository)

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
		todoRepository.StoreCalls(func(model *models.Todo) error {
			if model == nil {
				return &mockError{code: 500}
			}
			model.Id = positive.expect.Data.Id
			model.CreatedAt = positive.expect.Data.CreatedAt
			model.UpdatedAt = positive.expect.Data.UpdatedAt
			return nil
		})

		// action
		res := todoService.Store(positive.param)

		// assert response message
		assert.Equalf(
			t,
			positive.expect.Message,
			res.GetMessage(),
			"should return the correct response message\nMessage: %s\nExpect: %s",
			res.GetMessage(), positive.expect.Message,
		)

		// assert response data
		assert.Exactlyf(
			t,
			positive.expect.Data,
			res.Data,
			"should return the correct response data\nData: %#v\nExpect: %#v",
			res.Data, positive.expect.Data,
		)

		// assert response code
		assert.Equalf(
			t,
			res.GetCode(),
			201,
			"should return the correct response code\nCode: %s\nExpect: %s",
			res.GetCode(), 201,
		)

		// assert method calls count
		assert.Equalf(
			t,
			1,
			todoValidator.ValidateStoreCallCount(),
			"ValidateStore has to be called once\nCount: %d\nExpect: %d",
			todoValidator.ValidateDetailCallCount(), 1,
		)
		assert.Equalf(
			t,
			1,
			todoRepository.StoreCallCount(),
			"Store has to be called once\nCount: %d\nExpect: %d",
			todoRepository.StoreCallCount(), 1,
		)
	})

	t.Run("Test response code 400", func(t *testing.T) {
		// arrange
		todoValidator := validatorsfakes.FakeTodoValidator{}
		todoRepository := repositoriesfakes.FakeTodoRepository{}
		todoService := NewTodoService(&todoValidator, &todoRepository)

		negative := test{
			param: dto.StoreTodoRequest{},
			expect: entities.BaseResponse[dto.StoreTodoResponse]{
				Message: "some validation is failed",
				Data:    nil,
			},
		}

		// mock
		todoValidator.ValidateStoreReturns(fmt.Errorf("some validation is failed"))

		// action
		res := todoService.Store(negative.param)

		// assert response message
		assert.Equalf(
			t,
			negative.expect.Message,
			res.GetMessage(),
			"should have the correct response message\nMessage: %s\nExpect: %s",
			res.GetMessage(), negative.expect.Message,
		)

		// assert response data
		assert.Equalf(
			t,
			negative.expect.Data,
			res.Data,
			"response data should have the nil value\nData: %#v\nExpect: %#v",
			res.Data, negative.expect.Data,
		)

		// assert response code
		assert.Equalf(
			t,
			400,
			res.GetCode(),
			"should have the correct response code\nCode: %d\nExpect: %d",
			res.GetCode(), 400,
		)

		// assert method calls count
		assert.Equalf(
			t,
			todoValidator.ValidateStoreCallCount(),
			1,
			"ValidateStore has to be called once\nCount: %d\nExpect: %d",
			todoValidator.ValidateDetailCallCount(), 1,
		)
		assert.Equalf(
			t,
			todoRepository.StoreCallCount(),
			0,
			"StoreCallCount should no be called if request validation is failed\nCount: %d\nExpect: %d",
			todoRepository.StoreCallCount(), 0,
		)
	})

	t.Run("Test response 500", func(t *testing.T) {
		// arrange
		todoValidator := validatorsfakes.FakeTodoValidator{}
		todoRepository := repositoriesfakes.FakeTodoRepository{}
		todoService := NewTodoService(&todoValidator, &todoRepository)

		negative := test{
			param: dto.StoreTodoRequest{},
		}

		// mock
		wrapDBErr = func(err error) (code int) {
			val := err.(*mockError)
			return val.code
		}
		todoValidator.ValidateStoreReturns(nil)
		todoRepository.StoreReturns(&mockError{
			code: 500,
		})

		// action
		res := todoService.Store(negative.param)

		// assert response code
		assert.Equalf(
			t,
			res.GetCode(),
			500,
			"should return response code 500 if it has server error\nCode: %d\nExpect: %d",
			res.GetCode(), 500,
		)
	})
}
