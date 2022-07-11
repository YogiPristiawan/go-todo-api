package main

import (
	authApp "github.com/YogiPristiawan/go-todo-api/applications/auth"
	profileApp "github.com/YogiPristiawan/go-todo-api/applications/profile"
	todoApp "github.com/YogiPristiawan/go-todo-api/applications/todo"
	userApp "github.com/YogiPristiawan/go-todo-api/applications/user"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api"
	"github.com/YogiPristiawan/go-todo-api/modules/database"
	"github.com/YogiPristiawan/go-todo-api/modules/http"
	"github.com/YogiPristiawan/go-todo-api/modules/middleware"
	"github.com/YogiPristiawan/go-todo-api/modules/validator"
	"github.com/YogiPristiawan/go-todo-api/modules/validator/translate"
)

func main() {

	authMiddleware := middleware.CreateAuthMiddleware()

	mySqlDB := database.CreateMySqlConnection()

	// register validator
	validator := validator.CreateRequestValidator()
	validatorTranslator := translate.CreateRequestValidatorTranslate(validator)

	// register repository
	userRepository := userApp.NewUserRepository(mySqlDB)
	todoRepository := todoApp.NewTodoRepository(mySqlDB)

	// register use case
	userUseCase := userApp.NewUserUseCase(userRepository)
	authUseCase := authApp.NewAuthUseCase(userRepository)
	profileUseCase := profileApp.NewProfileUseCase(userRepository)
	todoUseCase := todoApp.NewTodoUseCase(todoRepository)

	http := http.CreateServer()

	// register routes
	api.CreateAuthRoute(map[string]any{
		"http":       http,
		"validator":  validator,
		"translator": validatorTranslator,
		"useCase":    authUseCase,
	})
	api.CreateUserRoute(map[string]any{
		"http":    http,
		"useCase": userUseCase,
	})
	api.CreateProfileRoute(map[string]any{
		"http":           http,
		"authMiddleware": authMiddleware,
		"useCase":        profileUseCase,
	})
	api.CreateTodoRoute(map[string]any{
		"http":           http,
		"authMiddleware": authMiddleware,
		"validator":      validator,
		"translator":     validatorTranslator,
		"useCase":        todoUseCase,
	})

	// init server
	http.Logger.Fatal(http.Start(":8080"))
}
