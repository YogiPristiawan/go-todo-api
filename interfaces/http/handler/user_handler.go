package handler

import (
	"github.com/YogiPristiawan/go-todo-api/domain/user"
	"github.com/YogiPristiawan/go-todo-api/modules/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UseCase user.UserUseCase
}

func NewUserHandler(useCase user.UserUseCase) *UserHandler {
	return &UserHandler{
		UseCase: useCase,
	}
}

func (u *UserHandler) Get(c echo.Context) error {
	// call use case
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
	// collect param
	id, err := helper.CollectParamUint(c, "id")
	if err != nil {
		return helper.HandleError(c, err)
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
