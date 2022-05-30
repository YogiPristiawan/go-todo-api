package use_case

import (
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type UserUseCase struct {
	UserRepository users.UserRepository
	Tokenize       *tokenize.JwtToken
	Validator      *validator.Validate
	Translator     ut.Translator
}

func (u *UserUseCase) GetAllUsers() ([]*entities.GetUsersResponse, error) {
	users, err := u.UserRepository.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return entities.MapGetUsersResponse(users), err
}

func (u *UserUseCase) DetailUserById(userId uint) (*entities.GetUserByIdResponse, error) {
	user, err := u.UserRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return entities.MapGetUserByIdResponse(user), err
}
