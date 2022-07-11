package todo

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/domain/todo"
	"github.com/YogiPristiawan/go-todo-api/modules/exceptions"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (t *TodoRepository) Store(todo *todo.TodoModel) (*todo.TodoModel, error) {
	if err := t.DB.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *TodoRepository) GetByUserId(userId uint) ([]*todo.TodoModel, error) {
	var todos []*todo.TodoModel
	if err := t.DB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todos, nil
		}
	}

	return todos, nil
}

func (t *TodoRepository) VerifyTodoAccess(userId uint, todoId uint) error {
	var todoModel *todo.TodoModel
	if err := t.DB.Model(todo.TodoModel{}).Select(
		"user_id",
		"is_finished",
		"date").Where("id = ?", todoId).First(&todoModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.NewNotFoundError("todo not found")
		}
	}

	if todoModel.UserId != userId {
		return exceptions.NewAuthorizationError("you don't have access to this todo")
	}

	return nil
}

func (t *TodoRepository) UpdateById(todoId uint, payload *todo.TodoModel) (*todo.TodoModel, error) {
	var todoModel *todo.TodoModel
	if err := t.DB.Where("id = ?", todoId).First(&todoModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("todo not found")
		}
	}

	t.DB.Model(&todoModel).Updates(payload)

	return todoModel, nil
}

func (t *TodoRepository) DeleteById(todoId uint) error {
	t.DB.Where("id = ?", todoId).Delete(&todo.TodoModel{})

	return nil
}
