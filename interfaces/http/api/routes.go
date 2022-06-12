package api

import (
	"net/http"

	"github.com/YogiPristiawan/go-todo-api/domains"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/auth"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/profile"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/todo"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/user"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
)

var server domains.Server
var middleware domains.Middleware

func CreateRoutes() {
	container.Resolve(&server)
	container.Resolve(&middleware)

	e := server.GetHttp()

	e.GET("/test", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Testing middleware")
	}, middleware.GetAuth())

	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Welcome to go-todo-api")
	})

	user.InitRoutes()
	auth.InitRoutes()
	profile.InitRoutes()
	todo.InitRoutes()
}
