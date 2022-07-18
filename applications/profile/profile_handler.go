package profile

import (
	"github.com/YogiPristiawan/go-todo-api/domain/profile"
	"github.com/YogiPristiawan/go-todo-api/modules/helper"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	UseCase profile.ProfileUseCase
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
