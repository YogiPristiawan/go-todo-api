package user

import (
	"errors"

	"go_todo_api/domain/user"
	"go_todo_api/modules/exceptions"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Store(user *user.UserModel) (*user.UserModel, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Get() ([]*user.UserModel, error) {
	var user []*user.UserModel
	err := u.DB.Find(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) FindById(id uint) (*user.UserModel, error) {
	var user *user.UserModel
	err := u.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("user not found")
		}
		panic(err)
	}

	return user, nil
}

func (u *UserRepository) FindByUsername(username string) (*user.UserModel, error) {
	var user *user.UserModel
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
	if err := u.DB.Model(&user.UserModel{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return err
	}

	if count != 0 {
		return exceptions.NewInvariantError("username telah digunakan")
	}

	return nil
}
