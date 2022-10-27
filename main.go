package main

import (
	authApp "go_todo_api/applications/auth"
	profileApp "go_todo_api/applications/profile"
	todoApp "go_todo_api/applications/todo"
	userApp "go_todo_api/applications/user"
	"go_todo_api/interfaces/http/handler"
	"go_todo_api/interfaces/http/routes"
	"go_todo_api/modules/database"
	"go_todo_api/modules/http"
	"go_todo_api/modules/middleware"
	"go_todo_api/modules/validator"
	"go_todo_api/modules/validator/translate"
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
