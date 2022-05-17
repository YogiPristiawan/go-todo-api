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

func (u *userUseCase) GetAllUsers() ([]*entities.GetUsersResponse, error) {
	users, err := u.userRepository.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return entities.MapGetUsersResponse(users), err
}

func (u *userUseCase) DetailUserById(userId int) (*entities.GetUserByIdResponse, error) {
	user, err := u.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return entities.MapGetUserByIdResponse(user), err
}
