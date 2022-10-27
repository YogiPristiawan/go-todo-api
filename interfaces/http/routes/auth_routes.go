package routes

import (
	"go_todo_api/interfaces/http/handler"

	"github.com/labstack/echo/v4"
)

func CreateAuthRoute(http *echo.Echo, handler *handler.AuthHandler) {
	g := http.Group("/auth")
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}
