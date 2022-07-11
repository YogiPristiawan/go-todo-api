package user

import (
	"strconv"

	"github.com/YogiPristiawan/go-todo-api/domain/user"
	"github.com/YogiPristiawan/go-todo-api/modules/exceptions"
	"github.com/YogiPristiawan/go-todo-api/modules/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UseCase user.UserUseCase
}

func (u *UserHandler) Get(c echo.Context) error {
	users, err := u.UseCase.Get()

	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(
		c,
		"success get all users",
		users,
	)
}

func (u *UserHandler) FindById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError("parameter harus berupa integer"))
	}

	user, err := u.UseCase.FindById(uint(id))

	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonHttpOk(
		c,
		"detail user",
		user,
	)
}
