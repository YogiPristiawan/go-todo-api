package use_case

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/auth"
	"github.com/YogiPristiawan/go-todo-api/domains/auth/entities"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	userEntities "github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
)

type authUseCase struct {
	userRepository users.UserRepository
	tokenize       *tokenize.JwtToken
}

func NewAuthUseCase(
	userRepository users.UserRepository,
	tokenize *tokenize.JwtToken,
) auth.AuthUseCase {
	return &authUseCase{
		userRepository: userRepository,
		tokenize:       tokenize,
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
	accessToken := a.tokenize.GenerateAccessToken(user.ID)

	return &entities.AuthLoginResponse{
		AccessToken: accessToken,
	}, nil
}

func (a *authUseCase) Register(payload *entities.AuthRegisterRequest) (*entities.AuthRegisterResponse, error) {
	// verify user doesn't exist
	if err := a.userRepository.VerifyAvailableUsername(payload.Username); err != nil {
		return nil, err
	}

	// store user
	user, err := a.userRepository.Store(&userEntities.UserModel{
		Username:  payload.Username,
		Password:  helpers.HashPassword(payload.Password),
		Gender:    payload.Gender,
		BirthDate: payload.BirthDate,
	})

	if err != nil {
		return nil, err
	}

	// genereate access token
	accessToken := a.tokenize.GenerateAccessToken(user.ID)

	return &entities.AuthRegisterResponse{
		AccessToken: accessToken,
	}, nil
}
