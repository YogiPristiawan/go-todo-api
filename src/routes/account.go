package routes

import (
	"go_todo_api/src/account"
	"go_todo_api/src/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateAccountRoute(r *gin.Engine, controller account.AccountController) {
	g := r.Group("v1")

	g.GET("/profile", middlewares.AuthMiddleware(), controller.GetProfile)
}
