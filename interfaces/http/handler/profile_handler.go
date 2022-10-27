package handler

import (
	"go_todo_api/domain/profile"
	"go_todo_api/modules/helper"

	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	UseCase profile.ProfileUseCase
}

func NewProfileHandler(useCase profile.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{
		UseCase: useCase,
	}
}

func (p *ProfileHandler) FindById(c echo.Context) error {
	// get authenticated user
	auth, err := helper.DecodeAuthJwtPayload(c)
	if err != nil {
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := p.UseCase.FindByUserId(auth.UserId)
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "user profile", result)
}
