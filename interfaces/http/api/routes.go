package api

import (
	"net/http"

	"github.com/YogiPristiawan/go-todo-api/applications/use_case"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/repository"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/users"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Welcome to go-todo-api")
	})

	userRepository := repository.NewUserRepository(db)
	userUseCase := use_case.NewUserUseCase(userRepository)
	users.InitRoutes(e, userUseCase)
}
