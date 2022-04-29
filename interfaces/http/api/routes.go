package api

import (
	"net/http"

	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/users"
	"github.com/labstack/echo/v4"
)

func CreateRoutes(e *echo.Echo) {
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Welcome to go-todo-api")
	})
	users.Users(e)
}
