package users

import "github.com/YogiPristiawan/go-todo-api/domains/users/entities"

type UserRepository interface {
	Store(*entities.UserModel) (*entities.UserModel, error)
	GetAllUsers() ([]*entities.UserModel, error)
	GetUserById(id uint) (*entities.UserModel, error)
	FindUserByUsername(username string) (*entities.UserModel, error)
	VerifyAvailableUsername(username string) error
}
