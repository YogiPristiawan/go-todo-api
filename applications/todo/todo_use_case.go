package todo

import (
	"go_todo_api/domain/todo"
)

type TodoUseCase struct {
	TodoRepository todo.TodoRepository
}

func NewTodoUseCase(todoRepository todo.TodoRepository) todo.TodoUseCase {
	return &TodoUseCase{
		TodoRepository: todoRepository,
	}
}

func (t *TodoUseCase) Store(payload *todo.StoreTodoRequest) (*todo.StoreTodoResponse, error) {
	result, err := t.TodoRepository.Store(&todo.TodoModel{
		UserId:     payload.UserId,
		Todo:       payload.Todo,
		Date:       payload.Date,
		IsFinished: payload.IsFinished,
	})
	if err != nil {
		return nil, err
	}

	return todo.MapStoreTodoResponse(result), nil
}

func (t *TodoUseCase) GetByUserId(userId uint) ([]*todo.GetTodosResponse, error) {
	result, err := t.TodoRepository.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	return todo.MapGetTodosResponse(result), nil
}

func (t *TodoUseCase) DetailById(userId uint, todoId uint) (*todo.DetailTodoResponse, error) {
	err := t.TodoRepository.VerifyTodoAccess(userId, todoId)
	if err != nil {
		return nil, err
	}

	result, err := t.TodoRepository.FindById(todoId)

	return todo.MapDetailTodoReponse(result), nil
}

func (t *TodoUseCase) UpdateById(
	authUserId uint,
	todoId uint,
	payload *todo.UpdateTodoRequest,
) (*todo.UpdateTodoResponse, error) {
	if err := t.TodoRepository.VerifyTodoAccess(authUserId, todoId); err != nil {
		return nil, err
	}

	result, err := t.TodoRepository.UpdateById(todoId, &todo.TodoModel{
		Todo:       payload.Todo,
		Date:       payload.Date,
		IsFinished: payload.IsFinished,
	})
	if err != nil {
		return nil, err
	}

	return todo.MapUpdateTodoResponse(result), nil
}

func (t *TodoUseCase) DeleteById(userId uint, todoId uint) error {
	if err := t.TodoRepository.VerifyTodoAccess(userId, todoId); err != nil {
		return err
	}

	return t.TodoRepository.DeleteById(todoId)
}
