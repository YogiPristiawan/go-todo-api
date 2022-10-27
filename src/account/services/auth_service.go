package services

import (
	"fmt"

	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/models"
	"go_todo_api/src/account/repositories"
	"go_todo_api/src/account/validators"
	"go_todo_api/src/shared/entities"
)

// AuthServices is an abstract that contains
// methods required to handle auth related bussiness logic
type AuthService interface {
	Login(in dto.LoginRequest) (out entities.BaseResponse[dto.LoginResponse])
	Register(in dto.RegisterRequest) (out entities.BaseResponse[dto.RegisterResponse])
}

// authService provides methods to handle auth related
// bussiness logic
type authService struct {
	validator   validators.AuthValidator
	accountRepo repositories.AccountRepository
}

// NewAuthService creates an instance of authService
func NewAuthService(
	validator validators.AuthValidator,
	accountRepo repositories.AccountRepository,
) AuthService {
	return &authService{
		validator:   validator,
		accountRepo: accountRepo,
	}
}

// Login handle login bussiness logic
func (a *authService) Login(in dto.LoginRequest) (out entities.BaseResponse[dto.LoginResponse]) {
	// validate payload
	if err := a.validator.ValidateLogin(in); err != nil {
		out.SetResponse(400, err)
		return
	}

	// get user by username if already exists
	user, err := a.accountRepo.GetByUsername(in.Username)
	switch wrapDBErr(err) {
	case 404:
		out.SetResponse(404, fmt.Errorf("username tidak ditemukan"))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	// compare password
	if err := comparePassword(user.Password, in.Password); err != nil {
		out.SetResponse(401, fmt.Errorf("password tidak sesuai"))
	}

	// generate access token
	accessToken := generateAccessToken(user.Id)

	out.Message = "login success"
	out.Data.AccessToken = accessToken
	return
}

// Register handle register bussiness logic
func (a *authService) Register(in dto.RegisterRequest) (out entities.BaseResponse[dto.RegisterResponse]) {
	// validate paylod
	if err := a.validator.ValidateRegister(in); err != nil {
		out.SetResponse(400, err)
		return
	}

	// create user account
	account := models.Account{
		Username:  in.Username,
		Password:  hashPassword(in.Password),
		Gender:    in.Gender,
		BirthDate: in.BirthDate,
	}
	err := a.accountRepo.Create(&account)
	switch wrapDBErr(err) {
	case 500:
		out.SetResponse(500, err)
		return
	}

	// generate access token
	accessToken := generateAccessToken(account.Id)

	out.Message = "register success"
	out.Data.AccessToken = accessToken
	return
}
