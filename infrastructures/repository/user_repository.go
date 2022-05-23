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

func (u *userRepository) Store(user *entities.UserModel) (*entities.UserModel, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
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
			return nil, exceptions.NewNotFoundError("user not found")
		}
		panic(err)
	}

	return user, nil
}

func (u *userRepository) FindUserByUsername(username string) (*entities.UserModel, error) {
	var user *entities.UserModel
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("user not found")
		}
		panic(err)
	}

	return user, nil
}

func (u *userRepository) VerifyAvailableUsername(username string) error {
	var count int64
	if err := u.db.Model(&entities.UserModel{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return err
	}

	if count != 0 {
		return exceptions.NewInvariantError("username telah digunakan")
	}

	return nil
}
