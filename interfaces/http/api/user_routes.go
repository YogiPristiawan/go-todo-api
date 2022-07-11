package api

import (
	userApp "github.com/YogiPristiawan/go-todo-api/applications/user"
	"github.com/YogiPristiawan/go-todo-api/domain/user"
	"github.com/labstack/echo/v4"
)

func CreateUserRoute(container map[string]any) {
	handler := &userApp.UserHandler{
		UseCase: container["useCase"].(user.UserUseCase),
	}

	http := container["http"].(*echo.Echo)

	g := http.Group("/users")
	g.GET("", handler.Get)
	g.GET("/:id", handler.FindById)
}
