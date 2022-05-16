package users

import (
	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	useCase users.UserUseCase
}

func NewUsersHandler(usecase users.UserUseCase) *usersHandler {
	return &usersHandler{
		useCase: usecase,
	}
}

func (u *usersHandler) GetAllUsers(c echo.Context) error {

	users := u.useCase.GetAllUsers()

	return helpers.ResponseJsonHttpOk(c, &helpers.ResponseContract{
		Message: "Success get all users",
		Status:  "succes",
		Data:    users,
	})
}
