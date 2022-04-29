package users

import (
	"net/http"

	"github.com/YogiPristiawan/go-todo-api/applications/use_case"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	useCase use_case.UserUseCase
}

func (u *UsersHandler) GetAllUsers(c echo.Context) error {
	users := u.useCase.GetAllUsers()

	return c.JSON(http.StatusOK, users)
}
