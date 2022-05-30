package auth

import (
	"github.com/YogiPristiawan/go-todo-api/domains"
	"github.com/YogiPristiawan/go-todo-api/domains/auth"
	"github.com/golobby/container/v3"
)

var server domains.Server
var authUseCase auth.AuthUseCase
var validatorInterface domains.Validator

func InitRoutes() {
	container.Resolve(&server)
	container.Resolve(&authUseCase)
	container.Resolve(&validatorInterface)

	handler := &authHandler{
		useCase:              authUseCase,
		validator:            validatorInterface.GetValidator(),
		validatorTranslation: validatorInterface.GetTranslator(),
	}

	e := server.GetHttp()

	g := e.Group("/auth")
	g.POST("/login", handler.login)
	g.POST("/register", handler.register)
}
