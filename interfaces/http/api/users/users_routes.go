package users

import (
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/labstack/echo/v4"
)

func Users(
	e *echo.Echo,
	useCase users.UserUseCase,
) {
	usersHandler := NewUsersHandler(useCase)
	users := e.Group("/users")

	users.GET("", usersHandler.GetAllUsers)
}
