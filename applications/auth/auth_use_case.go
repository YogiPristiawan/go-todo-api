package auth

import (
	"errors"

	"go_todo_api/domain/auth"
	"go_todo_api/domain/user"
	"go_todo_api/modules/helper"
)

type AuthUseCase struct {
	UserRepository user.UserRepository
}

func NewAuthUseCase(userRepository user.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		UserRepository: userRepository,
	}
}

func (a *AuthUseCase) Login(payload *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := a.UserRepository.FindByUsername(payload.Username)
	if err != nil {
		return nil, err
	}

	if err := helper.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, errors.New("password tidak sesuai")
	}

	// generate access token
	accessToken := helper.GenerateAccessToken(user.ID)

	return &auth.LoginResponse{
		AccessToken: accessToken,
	}, nil
}

func (a *AuthUseCase) Register(payload *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// verify user doesn't exist
	if err := a.UserRepository.VerifyAvailableUsername(payload.Username); err != nil {
		return nil, err
	}

	// store user
	user, err := a.UserRepository.Store(&user.UserModel{
		Username:  payload.Username,
		Password:  helper.HashPassword(payload.Password),
		Gender:    payload.Gender,
		BirthDate: payload.BirthDate,
	})

	if err != nil {
		return nil, err
	}

	// genereate access token
	accessToken := helper.GenerateAccessToken(user.ID)

	return &auth.RegisterResponse{
		AccessToken: accessToken,
	}, nil
}
