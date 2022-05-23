package auth

import (
	"github.com/YogiPristiawan/go-todo-api/domains/auth"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var handler *AuthHandler

func InitRoutes(
	e *echo.Echo,
	useCase auth.AuthUseCase,
	validator *validator.Validate,
	validatorTranslation ut.Translator,
) {
	g := e.Group("/auth")
	handler = NewAuthHandler(useCase, validator, validatorTranslation)

	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}
