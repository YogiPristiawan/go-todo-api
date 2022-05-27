package user

import (
	"strconv"

	"github.com/YogiPristiawan/go-todo-api/applications/exceptions"
	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	useCase users.UserUseCase
}

func NewUserHandler(useCase users.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.useCase.GetAllUsers()

	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonHttpOk(
		c,
		"success get all users",
		users,
	)
}

func (u *UserHandler) DetailUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewInvariantError("parameter harus berupa integer"))
	}

	user, err := u.useCase.DetailUserById(uint(id))

	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonHttpOk(
		c,
		"detail user",
		user,
	)
}
