package routes

import (
	"go_todo_api/src/account"

	"github.com/gin-gonic/gin"
)

// CreateAuthRoute initialize route
// for auth related operation
func CreateAuthRoute(r *gin.Engine, controller account.AuthController) {
	g := r.Group("/v1/auth")

	g.POST("/login", controller.Login)
	g.POST("/register", controller.Register)
}
