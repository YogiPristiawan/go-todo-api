package repository

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/applications/exceptions"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Store(user *entities.UserModel) (*entities.UserModel, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers() ([]*entities.UserModel, error) {
	var user []*entities.UserModel
	err := u.DB.Find(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserById(id uint) (*entities.UserModel, error) {
	var user *entities.UserModel
	err := u.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("user not found")
		}
		panic(err)
	}

	return user, nil
}

func (u *UserRepository) FindUserByUsername(username string) (*entities.UserModel, error) {
	var user *entities.UserModel
	err := u.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("user not found")
		}
		panic(err)
	}

	return user, nil
}

func (u *UserRepository) VerifyAvailableUsername(username string) error {
	var count int64
	if err := u.DB.Model(&entities.UserModel{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return err
	}

	if count != 0 {
		return exceptions.NewInvariantError("username telah digunakan")
	}

	return nil
}
