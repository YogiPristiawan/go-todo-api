package users

import "github.com/YogiPristiawan/go-todo-api/domains/users/entities"

type UserUseCase interface {
	GetAllUsers() []*entities.GetUsersResponse
}