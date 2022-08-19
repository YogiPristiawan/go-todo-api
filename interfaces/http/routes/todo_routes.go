package routes

import (
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/handler"
	"github.com/labstack/echo/v4"
)

func CreateTodoRoute(http *echo.Echo, handler *handler.TodoHandler, authMiddleware echo.MiddlewareFunc) {
	g := http.Group("/todos")
	g.POST("", handler.Store, authMiddleware)
	g.GET("", handler.GetByUserId, authMiddleware)
	g.GET("/:id", handler.DetailById, authMiddleware)
	g.PUT("/:id", handler.UpdateById, authMiddleware)
	g.DELETE("/:id", handler.DeleteById, authMiddleware)
}
