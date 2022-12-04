package services

import (
	"fmt"
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/repositories"
	"go_todo_api/src/shared/entities"
)

// AccountService is an abstract that contains
// methods to handle account related bussiness logic
type AccountService interface {
	GetProfile(in dto.ProfileRequest) (out entities.BaseResponse[dto.ProfileResponse])
}

// accountService is a struct that has methods to handle
// account related bussiness logic
type accountService struct {
	accountRepo repositories.AccountRepository
}

// NewAccountService creates an instance of accountService
func NewAccountService(accountRepo repositories.AccountRepository) AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}

// GetProfile handle an action to show authenticated
// user profile data
func (a *accountService) GetProfile(in dto.ProfileRequest) (out entities.BaseResponse[dto.ProfileResponse]) {
	profile, err := a.accountRepo.GetProfileByUserId(in.AuthUserId)

	switch wrapDBErr(err) {
	case 404:
		out.SetResponse(404, fmt.Errorf("profile not found"))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	out.Message = "success get user profile"
	out.Data = &dto.ProfileResponse{}
	mapGetProfileToResponse(out.Data, &profile)

	return
}
