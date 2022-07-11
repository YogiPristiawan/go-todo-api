package api

import (
	authApp "github.com/YogiPristiawan/go-todo-api/applications/auth"
	"github.com/YogiPristiawan/go-todo-api/domain/auth"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateAuthRoute(container map[string]any) {
	handler := &authApp.AuthHandler{
		UseCase:              container["useCase"].(auth.AuthUseCase),
		Validator:            container["validator"].(*validator.Validate),
		ValidatorTranslation: container["translator"].(ut.Translator),
	}

	http := container["http"].(*echo.Echo)

	g := http.Group("/auth")
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}
