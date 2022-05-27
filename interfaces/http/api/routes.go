package api

import (
	"net/http"

	"github.com/YogiPristiawan/go-todo-api/applications/use_case"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/repository"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/auth"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/profile"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api/user"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateRoutes(
	e *echo.Echo,
	db *gorm.DB,
	validator *validator.Validate,
	validatorTranslator ut.Translator,
	middleware echo.MiddlewareFunc,
	tokenize *tokenize.JwtToken,
) {
	e.GET("/test", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Testing middleware")
	}, middleware)

	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Welcome to go-todo-api")
	})

	// user route
	userRepository := repository.NewUserRepository(db)
	userUseCase := use_case.NewUserUseCase(userRepository, tokenize)
	user.InitRoutes(e, userUseCase)

	// auth route
	authUseCase := use_case.NewAuthUseCase(userRepository, tokenize)
	auth.InitRoutes(e, authUseCase, validator, validatorTranslator)

	// profile route
	profileUseCase := use_case.NewProfileUseCase(userRepository)
	profile.InitRoutes(e, profileUseCase, tokenize, middleware)
}
