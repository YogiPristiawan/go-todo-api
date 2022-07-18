package api

import (
	todoApp "github.com/YogiPristiawan/go-todo-api/applications/todo"
	"github.com/YogiPristiawan/go-todo-api/domain/todo"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateTodoRoute(container map[string]any) {
	handler := &todoApp.TodoHandler{
		UseCase:              container["useCase"].(todo.TodoUseCase),
		Validator:            container["validator"].(*validator.Validate),
		ValidatorTranslation: container["translator"].(ut.Translator),
	}

	authMiddleware := container["authMiddleware"].(echo.MiddlewareFunc)
	http := container["http"].(*echo.Echo)

	g := http.Group("/todos")
	g.POST("", handler.Store, authMiddleware)
	g.GET("", handler.GetByUserId, authMiddleware)
	g.GET("/:id", handler.DetailById, authMiddleware)
	g.PUT("/:id", handler.UpdateById, authMiddleware)
	g.DELETE("/:id", handler.DeleteById, authMiddleware)
}
