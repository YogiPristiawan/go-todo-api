package api

import (
	profileApp "github.com/YogiPristiawan/go-todo-api/applications/profile"
	"github.com/YogiPristiawan/go-todo-api/domain/profile"
	"github.com/labstack/echo/v4"
)

func CreateProfileRoute(container map[string]any) {
	handler := &profileApp.ProfileHandler{
		UseCase: container["useCase"].(profile.ProfileUseCase),
	}

	http := container["http"].(*echo.Echo)
	authMiddleware := container["authMiddleware"].(echo.MiddlewareFunc)

	g := http.Group("/profile", authMiddleware)

	g.GET("", handler.FindById)
}
