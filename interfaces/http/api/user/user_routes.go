package user

import (
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/labstack/echo/v4"
)

var userHandler *UserHandler

func InitRoutes(e *echo.Echo, useCase users.UserUseCase) {
	g := e.Group("/users")
	userHandler = NewUserHandler(useCase)

	g.GET("", userHandler.GetAllUsers)
	g.GET("/:id", userHandler.DetailUser)
}
