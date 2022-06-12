package todo

import (
	"github.com/YogiPristiawan/go-todo-api/domains"
	"github.com/YogiPristiawan/go-todo-api/domains/todo"
	"github.com/golobby/container/v3"
)

var server domains.Server
var todoUseCase todo.TodoUseCase
var validatorInterface domains.Validator
var middleware domains.Middleware

func InitRoutes() {
	container.Resolve(&server)
	container.Resolve(&todoUseCase)
	container.Resolve(&validatorInterface)
	container.Resolve(&middleware)

	handler := &todoHandler{
		useCase:              todoUseCase,
		validator:            validatorInterface.GetValidator(),
		validatorTranslation: validatorInterface.GetTranslator(),
	}

	e := server.GetHttp()

	g := e.Group("/todos")
	g.POST("", handler.storeTodo, middleware.GetAuth())
	g.GET("", handler.getTodos, middleware.GetAuth())
	g.PUT("/:id", handler.updateTodo, middleware.GetAuth())
	g.DELETE("/:id", handler.deleteTodo, middleware.GetAuth())
}
