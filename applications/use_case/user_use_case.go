package use_case

import (
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/domains/users/entities"
)

type userUseCase struct {
	userRepository users.UserRepository
}

func NewUserUseCase(r users.UserRepository) users.UserUseCase {
	return &userUseCase{
		userRepository: r,
	}
}

func (u *userUseCase) GetAllUsers() []*entities.GetUsersResponse {
	users := u.userRepository.GetAllUsers()
	return entities.MapGetUsersResponse(users)
}
