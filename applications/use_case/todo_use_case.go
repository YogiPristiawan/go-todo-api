package use_case

import (
	"github.com/YogiPristiawan/go-todo-api/domains/todo"
	"github.com/YogiPristiawan/go-todo-api/domains/todo/entities"
)

type TodoUseCase struct {
	TodoRepository todo.TodoRepository
}

func (t *TodoUseCase) StoreTodo(payload *entities.PostTodoRequest) (*entities.PostTodoResponse, error) {
	result, err := t.TodoRepository.StoreTodo(&entities.TodoModel{
		UserId:     payload.UserId,
		Todo:       payload.Todo,
		Date:       payload.Date,
		IsFinished: payload.IsFinished,
	})
	if err != nil {
		return nil, err
	}

	return entities.MapPostTodoResponse(result), nil
}

func (t *TodoUseCase) GetTodos(userId uint) ([]*entities.GetTodoResponse, error) {
	result, err := t.TodoRepository.GetTodosByUserId(userId)
	if err != nil {
		return nil, err
	}

	return entities.MapGetTodoResponse(result), nil
}

func (t *TodoUseCase) UpdateTodo(
	authUserId uint,
	todoId uint,
	payload *entities.PutTodoRequest,
) (*entities.PutTodoResponse, error) {
	if err := t.TodoRepository.VerifyTodoAccess(authUserId, todoId); err != nil {
		return nil, err
	}

	result, err := t.TodoRepository.UpdateById(todoId, &entities.TodoModel{
		Todo:       payload.Todo,
		Date:       payload.Date,
		IsFinished: payload.IsFinished,
	})
	if err != nil {
		return nil, err
	}

	return entities.MapPutTodoResponse(result), nil
}

func (t *TodoUseCase) DeleteTodo(userId uint, todoId uint) error {
	if err := t.TodoRepository.VerifyTodoAccess(userId, todoId); err != nil {
		return err
	}

	return t.TodoRepository.DeleteById(todoId)
}
