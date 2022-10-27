package routes

import (
	"go_todo_api/src/shared/middlewares"
	"go_todo_api/src/todo"

	"github.com/gin-gonic/gin"
)

func CreateTodoRoute(r *gin.Engine, controller todo.TodoController) {
	g := r.Group("v1")

	g.POST("/todos", middlewares.AuthMiddleware(), controller.Store)
	g.GET("/todos", middlewares.AuthMiddleware(), controller.Find)
	g.GET("/todos/:id", middlewares.AuthMiddleware(), controller.Detail)
	g.PUT("/todos/:id", middlewares.AuthMiddleware(), controller.Update)
	g.DELETE("/todos/:id", middlewares.AuthMiddleware(), controller.Delete)
}
