package use_case

import (
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
)

type userUseCase struct {
	userRepository users.UserRepository
	tokenize       *tokenize.JwtToken
}

func NewUserUseCase(r users.UserRepository, tokenize *tokenize.JwtToken) users.UserUseCase {
	return &userUseCase{
		userRepository: r,
		tokenize:       tokenize,
	}
}

func (u *userUseCase) GetAllUsers() ([]*entities.GetUsersResponse, error) {
	users, err := u.userRepository.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return entities.MapGetUsersResponse(users), err
}

func (u *userUseCase) DetailUserById(userId uint) (*entities.GetUserByIdResponse, error) {
	user, err := u.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return entities.MapGetUserByIdResponse(user), err
}
