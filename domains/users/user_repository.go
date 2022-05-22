package users

import "github.com/YogiPristiawan/go-todo-api/domains/users/entities"

type UserRepository interface {
	GetAllUsers() ([]*entities.UserModel, error)
	GetUserById(id int) (*entities.UserModel, error)
	FindUserByUsername(username string) (*entities.UserModel, error)
}
