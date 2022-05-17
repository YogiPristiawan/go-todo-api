package repository

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/applications/exceptions"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) GetAllUsers() ([]*entities.UserModel, error) {
	var user []*entities.UserModel
	err := u.db.Find(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("data not found")
		}
		panic(err)
	}

	return user, nil
}

func (u *userRepository) GetUserById(id int) (*entities.UserModel, error) {
	var user *entities.UserModel
	err := u.db.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("data not found")
		}
		panic(err)
	}

	return user, nil
}
