package main

import (
	authApp "github.com/YogiPristiawan/go-todo-api/applications/auth"
	profileApp "github.com/YogiPristiawan/go-todo-api/applications/profile"
	todoApp "github.com/YogiPristiawan/go-todo-api/applications/todo"
	userApp "github.com/YogiPristiawan/go-todo-api/applications/user"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/handler"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/routes"
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

	// register handler
	authHandler := handler.NewAuthHandler(authUseCase, validator, validatorTranslator)
	userHandler := handler.NewUserHandler(userUseCase)
	profileHandler := handler.NewProfileHandler(profileUseCase)
	todoHandler := handler.NewTodoHandler(todoUseCase, validator, validatorTranslator)

	http := http.CreateServer()

	// register routes
	routes.CreateAuthRoute(http, authHandler)
	routes.CreateUserRoute(http, userHandler, authMiddleware)
	routes.CreateProfileRoute(http, profileHandler, authMiddleware)
	routes.CreateTodoRoute(http, todoHandler, authMiddleware)

	// init server
	http.Logger.Fatal(http.Start(":8080"))
}
