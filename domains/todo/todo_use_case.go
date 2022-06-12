package todo

import (
	"github.com/YogiPristiawan/go-todo-api/domains/todo/entities"
)

type TodoUseCase interface {
	StoreTodo(payload *entities.PostTodoRequest) (*entities.PostTodoResponse, error)
	GetTodos(userId uint) ([]*entities.GetTodoResponse, error)
	UpdateTodo(authUserid uint, todoId uint, payload *entities.PutTodoRequest) (*entities.PutTodoResponse, error)
	DeleteTodo(authUserId uint, todoId uint) error
}
