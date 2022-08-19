package routes

import (
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/handler"
	"github.com/labstack/echo/v4"
)

func CreateProfileRoute(http *echo.Echo, handler *handler.ProfileHandler, authMiddleware echo.MiddlewareFunc) {
	g := http.Group("/profile", authMiddleware)

	g.GET("", handler.FindById)
}
