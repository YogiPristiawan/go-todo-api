package main

import (
	"fmt"
	"go_todo_api/src/routes"
	"go_todo_api/src/shared/databases"
	"go_todo_api/src/shared/validators"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"go_todo_api/src/account"
	accountRepo "go_todo_api/src/account/repositories"
	accountService "go_todo_api/src/account/services"
	accountValidator "go_todo_api/src/account/validators"

	"go_todo_api/src/todo"
	todoRepo "go_todo_api/src/todo/repositories"
	todoService "go_todo_api/src/todo/services"
	todoValidator "go_todo_api/src/todo/validators"
)

func main() {
	r := gin.Default()

	// INITIALIZE DATABASE
	db := databases.NewPgConn()
	defer db.Close()

	// INITIALIZE LIBRARIES
	validator := validators.NewValidatorAdapter()

	// INITIALIZE REPOSITORIES
	accountRepoImpl := accountRepo.NewAccountRepository(db)
	todoRepoImpl := todoRepo.NewTodoRepository(db)

	// INITIALIZE VALIDATORS
	authValidatorImpl := accountValidator.NewAuthValidator(validator)
	todoValidatorImpl := todoValidator.NewTodoValidator(validator)

	// INITIALIZE SERVICES
	authServiceImpl := accountService.NewAuthService(authValidatorImpl, accountRepoImpl)
	accountServiceImpl := accountService.NewAccountService(accountRepoImpl)
	todoServiceImpl := todoService.NewTodoService(todoValidatorImpl, todoRepoImpl)

	// INITIALIZE CONTROLLERS
	authControllerImpl := account.NewAuthController(authServiceImpl)
	accountControllerImpl := account.NewAccountController(accountServiceImpl)
	todoControllerImpl := todo.NewTodoController(todoServiceImpl)

	// INITIALIZE ROUTES
	routes.CreateAuthRoute(r, authControllerImpl)
	routes.CreateAccountRoute(r, accountControllerImpl)
	routes.CreateTodoRoute(r, todoControllerImpl)

	log.Fatal(r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))

}
