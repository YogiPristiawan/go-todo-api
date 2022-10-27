package routes

import (
	"go_todo_api/interfaces/http/handler"

	"github.com/labstack/echo/v4"
)

func CreateUserRoute(http *echo.Echo, handler *handler.UserHandler, authMiddleware echo.MiddlewareFunc) {
	g := http.Group("/users")
	g.GET("", handler.Get)
	g.GET("/:id", handler.FindById)
}
