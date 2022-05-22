package use_case

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/auth"
	"github.com/YogiPristiawan/go-todo-api/domains/auth/entities"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
)

type authUseCase struct {
	userRepository users.UserRepository
}

func NewAuthUseCase(userRepository users.UserRepository) auth.AuthUseCase {
	return &authUseCase{
		userRepository: userRepository,
	}
}

func (a *authUseCase) Login(payload *entities.AuthLoginRequest) (*entities.AuthLoginResponse, error) {
	user, err := a.userRepository.FindUserByUsername(payload.Username)
	if err != nil {
		return nil, err
	}

	if err := helpers.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, errors.New("password tidak sesuai")
	}

	// generate access token
	accessToken := helpers.GenerateAccessToken(user.ID)

	return &entities.AuthLoginResponse{
		AccessToken: accessToken,
	}, nil
}
