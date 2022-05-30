package use_case

import (
	"errors"

	"github.com/YogiPristiawan/go-todo-api/domains/auth/entities"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	userEntities "github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/encrypt"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
)

type AuthUseCase struct {
	UserRepository  users.UserRepository
	Tokenize        *tokenize.JwtToken
	PasswordManager *encrypt.HashPassword
}

func (a *AuthUseCase) Login(payload *entities.AuthLoginRequest) (*entities.AuthLoginResponse, error) {
	user, err := a.UserRepository.FindUserByUsername(payload.Username)
	if err != nil {
		return nil, err
	}

	if err := a.PasswordManager.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, errors.New("password tidak sesuai")
	}

	// generate access token
	accessToken := a.Tokenize.GenerateAccessToken(user.ID)

	return &entities.AuthLoginResponse{
		AccessToken: accessToken,
	}, nil
}

func (a *AuthUseCase) Register(payload *entities.AuthRegisterRequest) (*entities.AuthRegisterResponse, error) {
	// verify user doesn't exist
	if err := a.UserRepository.VerifyAvailableUsername(payload.Username); err != nil {
		return nil, err
	}

	// store user
	user, err := a.UserRepository.Store(&userEntities.UserModel{
		Username:  payload.Username,
		Password:  a.PasswordManager.HashPassword(payload.Password),
		Gender:    payload.Gender,
		BirthDate: payload.BirthDate,
	})

	if err != nil {
		return nil, err
	}

	// genereate access token
	accessToken := a.Tokenize.GenerateAccessToken(user.ID)

	return &entities.AuthRegisterResponse{
		AccessToken: accessToken,
	}, nil
}
