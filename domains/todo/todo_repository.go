package todo

import (
	"github.com/YogiPristiawan/go-todo-api/domains/todo/entities"
)

type TodoRepository interface {
	StoreTodo(*entities.TodoModel) (*entities.TodoModel, error)
	GetTodosByUserId(userId uint) ([]*entities.TodoModel, error)
	UpdateById(todoId uint, todo *entities.TodoModel) (*entities.TodoModel, error)
	DeleteById(todoId uint) error

	VerifyTodoAccess(userId uint, todoId uint) error
}
