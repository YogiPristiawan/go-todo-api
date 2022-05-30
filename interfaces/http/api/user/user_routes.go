package user

import (
	"github.com/YogiPristiawan/go-todo-api/domains"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/golobby/container/v3"
)

var server domains.Server
var userUseCase users.UserUseCase

func InitRoutes() {
	container.Resolve(&server)
	container.Resolve(&userUseCase)

	handler := &userHandler{
		useCase: userUseCase,
	}

	e := server.GetHttp()

	g := e.Group("/users")
	g.GET("", handler.getAllUsers)
	g.GET("/:id", handler.detailUser)
}
