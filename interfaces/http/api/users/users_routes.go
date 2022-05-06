package users

import (
	"github.com/YogiPristiawan/go-todo-api/applications/use_case"
	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo) {
	usersHandler := New(&use_case.UserUseCase{})
	users := e.Group("/users")

	users.GET("", usersHandler.GetAllUsers)
}
