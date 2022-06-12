package repository

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/applications/exceptions"
	"github.com/YogiPristiawan/go-todo-api/domains/todo/entities"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (t *TodoRepository) StoreTodo(todo *entities.TodoModel) (*entities.TodoModel, error) {
	if err := t.DB.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *TodoRepository) GetTodosByUserId(userId uint) ([]*entities.TodoModel, error) {
	var todos []*entities.TodoModel
	if err := t.DB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todos, nil
		}
	}

	return todos, nil
}

func (t *TodoRepository) VerifyTodoAccess(userId uint, todoId uint) error {
	var todo *entities.TodoModel
	if err := t.DB.Model(entities.TodoModel{}).Select(
		"user_id",
		"is_finished",
		"date").Where("id = ?", todoId).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.NewNotFoundError("todo not found")
		}
	}

	if todo.UserId != userId {
		return exceptions.NewAuthorizationError("you don't have access to this todo")
	}

	return nil
}

func (t *TodoRepository) UpdateById(todoId uint, payload *entities.TodoModel) (*entities.TodoModel, error) {
	var todo *entities.TodoModel
	if err := t.DB.Where("id = ?", todoId).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("todo not found")
		}
	}

	t.DB.Model(&todo).Updates(payload)

	return todo, nil
}

func (t *TodoRepository) DeleteById(todoId uint) error {
	t.DB.Where("id = ?", todoId).Delete(&entities.TodoModel{})

	return nil
}
