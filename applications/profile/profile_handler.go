package profile

import (
	"strings"

	"github.com/YogiPristiawan/go-todo-api/domain/profile"
	"github.com/YogiPristiawan/go-todo-api/modules/helper"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	UseCase profile.ProfileUseCase
}

func (p *ProfileHandler) FindById(c echo.Context) error {
	authorization := c.Request().Header["Authorization"]
	token := strings.Split(authorization[0], " ")[1]

	claims, _ := helper.DecodeAccessToken(token)

	result, err := p.UseCase.FindByUserId(claims.UserId)
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(c, "user profile", result)
}
