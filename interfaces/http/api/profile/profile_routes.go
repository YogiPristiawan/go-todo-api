package profile

import (
	"github.com/YogiPristiawan/go-todo-api/domains/profile"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	"github.com/labstack/echo/v4"
)

var handler *ProfileHandler

func InitRoutes(
	e *echo.Echo,
	useCase profile.ProfileUseCase,
	tokenize *tokenize.JwtToken,
	middleware echo.MiddlewareFunc,
) {
	handler = NewProfileHandler(useCase, tokenize)
	g := e.Group("/profile", middleware)

	g.GET("", handler.GetProfile)
}
